package main

import (
	"fmt"
	"log"
	"strings"

	apiclient "bigcommerce.com/apis/clients/wishlists/client"
	"bigcommerce.com/apis/clients/wishlists/client/wishlists"
	httptransport "github.com/go-openapi/runtime/client"
)

func getWishlists() {
	config := getConfig()
	apiKeyHeaderAuth := httptransport.APIKeyAuth("X-Auth-Token", "header", fmt.Sprint(config["accessToken"]))

	basePath := strings.Replace(apiclient.DefaultTransportConfig().BasePath, "{$$.env.store_hash}", fmt.Sprint(config["storeId"]), 1)
	transportConfig := apiclient.DefaultTransportConfig().WithBasePath(basePath)
	client := apiclient.NewHTTPClientWithConfig(nil, transportConfig)
	params := wishlists.NewWishlistsGetParams()
	customerId := int32(999)
	params.SetCustomerID(&customerId)
	resp, err := client.Wishlists.WishlistsGet(params, apiKeyHeaderAuth)

	if err != nil {
		log.Print("Had an error")
		log.Fatal(err)
	}
	log.Printf("Payload: %#v\n", resp.Payload)
}
