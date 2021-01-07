package main

import (
	"fmt"
	"log"
	"strings"

	apiclient "bigcommerce.com/apis/clients/ordersv2/client"
	"bigcommerce.com/apis/clients/ordersv2/client/orders"
	httptransport "github.com/go-openapi/runtime/client"
)

func getOrders() {
	config := getConfig()
	apiKeyHeaderAuth := httptransport.APIKeyAuth("X-Auth-Token", "header", fmt.Sprint(config["accessToken"]))

	basePath := strings.Replace(apiclient.DefaultTransportConfig().BasePath, "{$$.env.store_hash}", fmt.Sprint(config["storeId"]), 1)
	log.Print(basePath)
	transportConfig := apiclient.DefaultTransportConfig().WithBasePath(basePath)
	client := apiclient.NewHTTPClientWithConfig(nil, transportConfig)
	resp, err := client.Orders.GetAllOrders(orders.NewGetAllOrdersParams(), apiKeyHeaderAuth)

	if err != nil {
		// get a 204 when there are no orders, but that wasn't in the spec, so we get an error
		log.Printf("Err (expected): %#v\n", err)
		log.Printf("Payload: %#v\n", resp)
	}
}
