package main

import "fmt"

func main() {
	fmt.Println("Maps in Go lang")

	// Creating Map

	m := make(map[string]string)

	// Setting element to the map

	m["name"] = "Mritunjay Paswan"

	// Accessing map elements

	fmt.Println(m["name"]) // If key does not exists in the map then it will return zeroed values.

	// Getting length in map

	length := len(m)
	fmt.Println(length)

	// Deleting key from the map
	delete(m, "name") // name is the key

	clear(m) // Clear function is used to delete / clear all the map data

	// Declaring map without make keyword

	mp := map[string]int{"phone": 30, "price": 5000}

	fmt.Println(mp)

	// Checking element is present in the map or not

	val, ok := mp["prices"] // val will contain the value if present otherwise it will contain the zeroed values, ok will contain the bool value true if element exists in map otherwise it returns false.

	fmt.Println(val, ok)
}
