package main

import "time"

type Customer struct {
	name        string
	phoneNumber string
}

// embedding struct Customer inside Order1
type Order1 struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	Customer
}

func hi() {
	order := Order1{
		id:     "54321",
		amount: 49.99,
		status: "processing",
		Customer: Customer{
			name:        "John Doe",
			phoneNumber: "123-456-7890",
		},
	}
	order.createdAt = time.Now()
	println(order.id)
	println(order.Customer.name)
	println(order.Customer.phoneNumber)
	println(order.createdAt.String())
}
