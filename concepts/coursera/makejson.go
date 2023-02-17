/*
Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and
add the name and address to the map using the keys “name” and “address”, respectively. Your program should use Marshal()
to create a JSON object from the map, and then your program should print the JSON object.
*/

package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var input_name, input_address string
	var person_map = make(map[string]string)

	fmt.Println("Please, enter a name: ")
	fmt.Scan(&input_name)

	fmt.Println("Please, enter an address: ")
	fmt.Scan(&input_address)

	person_map["name"] = input_name
	person_map["address"] = input_address

	person_json, err := json.Marshal(person_map)
	if err != nil {
		fmt.Println("Error in serializing to JSON: ", err)
	} else {
		fmt.Println(string(person_json))
	}
}
