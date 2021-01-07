package main

import (
	"fmt"
	"log"
	"strings"

	apiclient "bigcommerce.com/apis/clients/channels/client"
	"bigcommerce.com/apis/clients/channels/client/channels"
	httptransport "github.com/go-openapi/runtime/client"
)

func getChannels() {
	config := getConfig()
	apiKeyHeaderAuth := httptransport.APIKeyAuth("X-Auth-Token", "header", fmt.Sprint(config["accessToken"]))

	basePath := strings.Replace(apiclient.DefaultTransportConfig().BasePath, "{$$.env.store_hash}", fmt.Sprint(config["storeId"]), 1)
	transportConfig := apiclient.DefaultTransportConfig().WithBasePath(basePath)
	client := apiclient.NewHTTPClientWithConfig(nil, transportConfig)
	resp, err := client.Channels.ListChannels(channels.NewListChannelsParams(), apiKeyHeaderAuth)
	if err != nil {
		log.Print("Had an error")
		log.Fatal(err)
	}
	log.Printf("Payload: %#v\n", resp.Payload, resp.GetPayload())
}
