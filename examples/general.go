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

	news, err := client.GetNews(ctx, models.NewsQuery{From: "2025-02-06"})
	if err != nil {
		log.Fatal(err)
	}

	out, err := json.MarshalIndent(news, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
