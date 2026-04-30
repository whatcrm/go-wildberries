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

	warehouses, err := client.GetFBWWarehouses(ctx)
	if err != nil {
		log.Fatal(err)
	}
	printJSON("fbw warehouses", warehouses)

	tariffs, err := client.GetFBWTransitTariffs(ctx)
	if err != nil {
		log.Fatal(err)
	}
	printJSON("fbw transit tariffs", tariffs)

	supplies, err := client.GetFBWSupplies(ctx, models.FBWSuppliesQuery{
		Limit:  50,
		Offset: 0,
	}, models.FBWSuppliesFiltersRequest{
		StatusIDs: []models.FBWStatusID{
			models.FBWSupplyStatusAccepted,
			models.FBWSupplyStatusUnloadedOnGate,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	printJSON("fbw supplies", supplies)

	if len(supplies) == 0 {
		fmt.Println("no supplies returned, stop demo")
		return
	}

	supplyID := supplies[0].SupplyID
	if supplyID == nil {
		fmt.Println("first item is preorder, skip details demo")
		return
	}

	details, err := client.GetFBWSupplyDetails(ctx, *supplyID, models.FBWSupplyDetailsQuery{})
	if err != nil {
		log.Fatal(err)
	}
	printJSON("fbw supply details", details)

	goods, err := client.GetFBWSupplyGoods(ctx, *supplyID, models.FBWSupplyGoodsQuery{
		Limit:  100,
		Offset: 0,
	})
	if err != nil {
		log.Fatal(err)
	}
	printJSON("fbw supply goods", goods)
}

func printJSON(title string, v interface{}) {
	out, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("=== %s ===\n%s\n\n", title, string(out))
}
