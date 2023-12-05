package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World! from map/main.go")

	// Declare a map and initialize it
	myMapInit := make(map[string]int)
	myMapInit["dis"] = 1
	fmt.Println(myMapInit)

	// intialize a map with key-value pairs
	myMap := map[string]int{"one":1, "two":2, "three":3}
	fmt.Println(myMap)

	// Access a value from the map
	value := myMap["one"]
	fmt.Println("Value is: ",value)

	// Add key-value pairs to the map
	myMap["four"] = 4
	myMap["five"] = 5
	myMap["two"] = 22

	fmt.Println("After adding key-value/updating pairs: ", myMap)

	// Delete a key-value pair from the map
	delete(myMap, "three")
	fmt.Println("After deleting key-value pair: ", myMap)

	// Check if a key exists in the map
	// If the key exists, then the value is returned
	value,exist := myMap["two"]
	if exist {
		fmt.Println("Existing value is: ",value)
	} else {
		fmt.Println("Key does not exist")
	}

	// be rewritter as
	// Valid only in if-else block
	if value,exit := myMap["two"]; exit{
		fmt.Println("Existing value is: ",value)
	} else {
		fmt.Println("Key does not exist")
	}

	// iterate over the map
	for key,value := range myMap {
		fmt.Println("Key: ",key," Value: ",value)	
	}

	// length of the map
	length := len(myMap)
	fmt.Println("Length of the map is: ",length)

	// map of maps
	mapOfMaps := map[string]map[string]int{
		//"one": map[string]int{val} => is redundant
		"one" : {
			"i":1,
			"واحد":1,
		},
		"two" : {
			"ii":2,
		"اثنان":2,
		},
	}
	fmt.Println("Map of maps: ",mapOfMaps)


	// nil map
	var nilMap map[string]int; 
	fmt.Println("Nil map: ",nilMap)



	





}