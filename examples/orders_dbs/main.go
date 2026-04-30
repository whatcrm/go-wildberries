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

	// 1) Get new DBS orders.
	newOrders, err := client.GetDBSNewOrders(ctx)
	if err != nil {
		log.Fatal(err)
	}
	printJSON("new dbs orders", newOrders)

	// 2) Collect first IDs from response.
	orderIDs := extractOrderIDs(newOrders.Orders, 5)
	if len(orderIDs) == 0 {
		fmt.Println("no orders returned, stop demo")
		return
	}

	// 3) Read statuses.
	statuses, err := client.GetDBSStatusesInfo(ctx, models.DBSOrdersRequestV2{
		OrdersIDs: orderIDs,
	})
	if err != nil {
		log.Fatal(err)
	}
	printJSON("dbs statuses", statuses)

	// 4) Read customer info.
	clientInfo, err := client.GetDBSClientInfo(ctx, models.OrdersIDsRequest{
		Orders: orderIDs,
	})
	if err != nil {
		log.Fatal(err)
	}
	printJSON("dbs client info", clientInfo)

	// 5) Read metadata for the same orders.
	meta, err := client.GetDBSMetaInfo(ctx, models.DBSOrdersRequestV2{
		OrdersIDs: orderIDs,
	})
	if err != nil {
		log.Fatal(err)
	}
	printJSON("dbs meta", meta)
}

func extractOrderIDs(rawOrders []map[string]interface{}, max int) []int64 {
	result := make([]int64, 0, max)
	for _, o := range rawOrders {
		v, ok := o["id"]
		if !ok {
			continue
		}
		idFloat, ok := v.(float64)
		if !ok {
			continue
		}
		result = append(result, int64(idFloat))
		if len(result) >= max {
			break
		}
	}
	return result
}

func printJSON(title string, v interface{}) {
	out, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("=== %s ===\n%s\n\n", title, string(out))
}
