package main

import (
	"fmt"
	"reflect"
)

func sum(nums ...int) {
	fmt.Println(nums, reflect.TypeOf(nums))
}

func main() {
	fmt.Println("Welcome to variadic function")
	nums := []int{5, 6, 7, 8}
	sum(nums...)

}
