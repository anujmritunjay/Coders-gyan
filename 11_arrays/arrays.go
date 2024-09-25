package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays tutorial in Go lang")

	// var nums [4]int
	// // By default a zeroed values assinged to array element
	// // For int -> 0, bool -> false, string -> ''
	// fmt.Println(nums)

	// nums[1] = 8

	// fmt.Println(nums)

	// Short hand syntax

	nums := [4]int{1, 3}

	fmt.Println(nums)

}
