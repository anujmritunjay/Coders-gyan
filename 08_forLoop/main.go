package main

import "fmt"

// There is only one loop in Go that is for loop

func main() {
	fmt.Println("This is the for loop tutorial in Go lang.")

	// While loop

	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// Infinite loop

	// for {
	// 	fmt.Println("This loop is infinite")
	// }

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}
