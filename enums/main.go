package main

import "fmt"

// in go we implement enum using const
type OrderStatus int
type OrderStatus1 string

const (
	Pending   OrderStatus1 = "Pending"
	Completed OrderStatus1 = "Completed"
	Canceled  OrderStatus1 = "Canceled"
)

const (
	Recieved OrderStatus = iota
	Processing
	Shipped
	Delivered
)

func ChangeStatus(status OrderStatus) {
	fmt.Println("Order status changed to:", status)
}

func ChangeStatus1(status OrderStatus1) {
	fmt.Println("Order status changed to:", status)
}

func main() {
	ChangeStatus(Recieved)
	ChangeStatus(Processing)
	ChangeStatus(Shipped)
	ChangeStatus(Delivered) // output will be 0,1,2,3 because they are iota values

	ChangeStatus1(Pending) // output will be Pending, Completed, Canceled
}
