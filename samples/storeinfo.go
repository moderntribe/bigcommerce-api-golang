package main

import (
  "fmt"
  "log"
  "strings"

  apiclient "bigcommerce.com/apis/clients/storeinfo/client"
  "bigcommerce.com/apis/clients/storeinfo/client/store_information"
  httptransport "github.com/go-openapi/runtime/client"
)

func getStoreInfo() {
  // NB we have to
  // - fix the sitewidehttps_enabled which is a bool, not a string
  // - add an actual logo image to the store (otherwise api returns an empty array - bug)
  config := getConfig()
  apiKeyHeaderAuth := httptransport.APIKeyAuth("X-Auth-Token", "header", fmt.Sprint(config["accessToken"]))

  basePath := strings.Replace(apiclient.DefaultTransportConfig().BasePath, "{$$.env.store_hash}", fmt.Sprint(config["storeId"]), 1)
  transportConfig := apiclient.DefaultTransportConfig().WithBasePath(basePath)
  client := apiclient.NewHTTPClientWithConfig(nil, transportConfig)
  resp, err := client.StoreInformation.GetStore(store_information.NewGetStoreParams(), apiKeyHeaderAuth)
  if err != nil {
    log.Print("Had an error")
    log.Fatal(err)
  }
  log.Printf("Payload: %#v\n", resp.Payload)
}
