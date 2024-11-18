package main

import "fmt"

func main() {
	fmt.Println("Welcome to range tutorial")

	// Range function for the slice

	nums := []int{1, 2, 3, 4, 6}

	for _, val := range nums {
		fmt.Println(val)
	}

	// Range for the map

	// mp := map[string]string{"name": "Mritunjay Paswan", "Address": "Kusinagar"}

	// for k, v := range mp {
	// 	fmt.Println(k, v)
	// }
}
