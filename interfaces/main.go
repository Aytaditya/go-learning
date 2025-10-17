package main

import "fmt"

type Payment struct {
	gateway Stripe
}

func (p Payment) makePayment(amount float32) {
	// razorPayGateway := Razorpay{}
	// razorPayGateway.makePayment(amount)
	// stripeGateway := Stripe{}
	// stripeGateway.makePayment(amount)
	p.gateway.makePayment(amount)
}

type Razorpay struct {
}

func (r Razorpay) makePayment(amount float32) {
	// logic to make payment using Razorpay
	fmt.Println("Payment of", amount, "made using Razorpay")
}

type Stripe struct {
}

func (s Stripe) makePayment(amount float32) {
	// logic to make payment using Stripe
	fmt.Println("Payment of", amount, "made using Stripe")
}

func main() {
	pay := Payment{}
	pay.makePayment(1000)
}
