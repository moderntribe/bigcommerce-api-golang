package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	getThemes()
	getWidgets()
	getSubscribers()
	getScripts()
}
