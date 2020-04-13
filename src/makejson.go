package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {

	inputMap := map[string]string{"name": "", "address": ""}
	var input string

	fmt.Println("Enter a name: ")
	fmt.Scan(&input)
	inputMap["name"] = strings.Trim(input, " \n")
	fmt.Println("Enter an address: ")
	fmt.Scan(&input)
	inputMap["address"] = strings.Trim(input, " \n")

	jsonData, err := json.MarshalIndent(inputMap, "", "\t")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("\n%s\n", jsonData)
}
