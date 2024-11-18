package main

import "fmt"

func change(num *int) {
	*num = 5
}

func main() {
	fmt.Println("Welcome to the Pointers Tutorials")
	num := 1
	change(&num)
	fmt.Println(num)
}
