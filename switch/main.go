package main

import (
	"fmt"
	"time"
)

func main() {

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println(time.Now().Weekday())
		fmt.Println("It's the weekend")
	default:
		fmt.Println(time.Now().Weekday())
		fmt.Println("It's a weekday")
	}

	// varaibles can even stire functions in go lang
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("I am an integer")
		case string:
			fmt.Println("I am a string")
		default:
			fmt.Println("I am something else")
		}
	}
	whoAmI(5)
}
