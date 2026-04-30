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

	newOrders, err := client.GetClickCollectNewOrders(ctx)
	if err != nil {
		log.Fatal(err)
	}
	printJSON("click-collect new orders", newOrders)

	if len(newOrders.Orders) == 0 {
		return
	}

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
	if len(orderIDs) == 0 {
		return
	}

	statuses, err := client.GetClickCollectStatusesInfo(ctx, models.ClickCollectOrdersRequestV2{
		OrdersIDs: orderIDs,
	})
	if err != nil {
		log.Fatal(err)
	}
	printJSON("click-collect statuses", statuses)
}

func printJSON(title string, v interface{}) {
	out, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("=== %s ===\n%s\n\n", title, string(out))
}
