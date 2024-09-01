package main

import "fmt"

func main() {
	fmt.Println("This is the if else tutorial in Go lang")

	var age int = 17

	if age >= 18 {
		fmt.Println("Person is Adult")
	} else {
		fmt.Println("Person is not Adult")
	}
}
