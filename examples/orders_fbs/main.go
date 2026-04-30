package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	gowildberries "github.com/whatcrm/go-wildberries"
	"github.com/whatcrm/go-wildberries/models"
)

func main() {
	token := "WB_API_TOKEN"

	client, err := gowildberries.NewClient(token)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// 1) Get new FBS assembly orders.
	newOrders, err := client.GetNewOrders(ctx)
	if err != nil {
		log.Fatal(err)
	}
	printJSON("new orders", newOrders)

	// 2) Fetch statuses for first orders.
	if len(newOrders.Orders) > 0 {
		orderIDs := make([]int64, 0, len(newOrders.Orders))
		for _, rawOrder := range newOrders.Orders {
			if idValue, ok := rawOrder["id"]; ok {
				if idFloat, ok := idValue.(float64); ok {
					orderIDs = append(orderIDs, int64(idFloat))
				}
			}
			if len(orderIDs) == 5 {
				break
			}
		}

		if len(orderIDs) > 0 {
			statuses, err := client.GetOrdersStatus(ctx, models.OrdersIDsRequest{Orders: orderIDs})
			if err != nil {
				log.Fatal(err)
			}
			printJSON("orders status", statuses)
		}
	}

	// 3) Create a supply and list it.
	supply, err := client.CreateSupply(ctx, models.CreateSupplyRequest{
		Name: "SDK example supply",
	})
	if err != nil {
		log.Fatal(err)
	}
	printJSON("created supply", supply)

	// 4) Get supply details.
	supplyInfo, err := client.GetSupply(ctx, supply.ID)
	if err != nil {
		log.Fatal(err)
	}
	printJSON("supply info", supplyInfo)
}

func printJSON(title string, v interface{}) {
	out, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("=== %s ===\n%s\n\n", title, string(out))
}
