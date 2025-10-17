package main

import (
	"fmt"
	"time"
)

// structs acts as a blueprint for creating complex data types
type Order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

func main() {
	order := Order{
		id:     "12345",
		amount: 99.99,
		status: "pending",
	}
	order.createdAt = time.Now()
	println(order.id)
	fmt.Println(order.createdAt)
}
