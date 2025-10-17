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

// method with value receiver
func (o *Order) changeStatus(status string) {
	o.status = status
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

	order.changeStatus("shipped") // won't change the status of original order status if * is not used in method receiver
	fmt.Println("Order status:", order.status)

	// one time struct
	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println("Language:", language.name, "Is Good?", language.isGood)

}
