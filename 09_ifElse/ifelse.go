package main

import "fmt"

func main() {
	fmt.Println("Welcome to the if else tutorial in the Go lang")

	var age int = 17

	if age >= 18 {
		fmt.Println("Person is Adult")
	} else {
		fmt.Println("Person is not Adult")
	}
}
