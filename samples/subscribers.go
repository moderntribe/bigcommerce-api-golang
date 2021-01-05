package main

import (
	"fmt"
	"log"
	"strings"

	apiclient "bigcommerce.com/apis/clients/subscribers/client"
	"bigcommerce.com/apis/clients/subscribers/client/subscribers"
	httptransport "github.com/go-openapi/runtime/client"
)

func getSubscribers() {
	config := getConfig()
	apiKeyHeaderAuth := httptransport.APIKeyAuth("X-Auth-Token", "header", fmt.Sprint(config["accessToken"]))

	basePath := strings.Replace(apiclient.DefaultTransportConfig().BasePath, "{$$.env.store_hash}", fmt.Sprint(config["storeId"]), 1)
	transportConfig := apiclient.DefaultTransportConfig().WithBasePath(basePath)
	client := apiclient.NewHTTPClientWithConfig(nil, transportConfig)
	resp, err := client.Subscribers.GetSubscribers(subscribers.NewGetSubscribersParams(), apiKeyHeaderAuth)
	if err != nil {
		log.Print("Had an error")
		log.Fatal(err)
	}
	log.Printf("Payload: %#v\n", resp.Payload)
}
