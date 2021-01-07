package main

import (
	"fmt"
	"log"
	"strings"

	apiclient "bigcommerce.com/apis/clients/ordersv3/client"
	"bigcommerce.com/apis/clients/ordersv3/client/transactions"
	httptransport "github.com/go-openapi/runtime/client"
)

func getTransactions() {
	config := getConfig()
	apiKeyHeaderAuth := httptransport.APIKeyAuth("X-Auth-Token", "header", fmt.Sprint(config["accessToken"]))

	basePath := strings.Replace(apiclient.DefaultTransportConfig().BasePath, "{$$.env.store_hash}", fmt.Sprint(config["storeId"]), 1)
	transportConfig := apiclient.DefaultTransportConfig().WithBasePath(basePath)
	client := apiclient.NewHTTPClientWithConfig(nil, transportConfig)
	params := transactions.NewGetTransactionsParams()
	params.OrderID = 10
	ok, content, err := client.Transactions.GetTransactions(params, apiKeyHeaderAuth)
	log.Printf("OK: %#v\n", ok, ok)
	if err != nil {
		log.Printf("err (expected): %#v\n", err, err)
	}
	log.Printf("Content: %#v\n", content, content)
}
