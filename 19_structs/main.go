package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float64
	status    string
	createdAt time.Time // Nano second precision
}

func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getAmount() float64 {
	return o.amount
}

func main() {
	fmt.Println("Welcome to struct tutorial")
	myOrder := order{
		id:     "1",
		amount: 13.02,
		status: "received",
	}
	myOrder.createdAt = time.Now()

	myOrder.changeStatus("delivered")
	fmt.Println(myOrder)
	fmt.Println(myOrder.getAmount())
}
