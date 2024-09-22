package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to switch tutorial in Go lang")
	// i := 2

	// Simple switch
	// switch i {
	// case 1:
	// 	fmt.Println("One")

	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// default:
	// 	fmt.Println("Invalid choice")
	// }

	// Multiple condition switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("This is the weekend")
	default:
		fmt.Println("Its working Day")
	}
}
