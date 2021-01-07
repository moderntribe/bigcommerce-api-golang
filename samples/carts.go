package main

import (
	"fmt"
	"log"
	"strings"

	apiclient "bigcommerce.com/apis/clients/carts/client"
	"bigcommerce.com/apis/clients/carts/client/cart"
	httptransport "github.com/go-openapi/runtime/client"
)

func getCarts() {
	config := getConfig()
	apiKeyHeaderAuth := httptransport.APIKeyAuth("X-Auth-Token", "header", fmt.Sprint(config["accessToken"]))

	basePath := strings.Replace(apiclient.DefaultTransportConfig().BasePath, "{$$.env.store_hash}", fmt.Sprint(config["storeId"]), 1)
	transportConfig := apiclient.DefaultTransportConfig().WithBasePath(basePath)
	client := apiclient.NewHTTPClientWithConfig(nil, transportConfig)
	params := cart.NewGetACartParams()
	params.CartID = "NoExist"
	resp, err := client.Cart.GetACart(params, apiKeyHeaderAuth)
	if err != nil {
		log.Printf("Err (expected): %#v\n", err)
	}
	log.Printf("Response: %#v\n", resp)
}
