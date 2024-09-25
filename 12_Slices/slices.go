package main

import "fmt"

func main() {
	fmt.Println("Welcome to slices in Go lang")

	// Declare

	// var nums []int // By default it is equals to nil

	// fmt.Println(nums == nil)

	// fmt.Println(len(nums))

	var nums = make([]int, 0)

	// Adding element to the slices

	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5)
	nums = append(nums, 6)

	// fmt.Println(nums)

	// Copy slices

	// nums2 := make([]int, len(nums))
	// copy(nums2, nums)

	// fmt.Println(nums, nums2)

	// Slicing the slice

	// newSlice := nums[0:3]
	// newSlice := nums[:3] // Ommetig first values is will be consider 0 by default

	newSlice := nums[3:] // Ommeting last value will be consider for all values
	fmt.Println(newSlice)

}
