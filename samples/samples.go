package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func getConfig() map[string]interface{} {
	jsonFile, err := os.Open("../gulpfile.config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal(byteValue, &result)

	return result
}

func main() {
	getStoreInfo()
	log.Println("---")

	getThemes()
	log.Println("---")

	getWidgets()
	log.Println("---")

	getSubscribers()
	log.Println("---")

	getScripts()
	log.Println("---")

	getPricelists()
	log.Println("---")

	getOrders()
	log.Println("---")

	getTransactions()
	log.Println("---")

	getChannels()
	log.Println("---")

	getCarts()
	log.Println("---")

	getWishlists()
	log.Println("---")

	getSites()
}
