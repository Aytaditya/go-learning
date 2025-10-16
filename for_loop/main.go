package main

import "fmt"

func main() {
	// implementing while loop using for loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// infinte loop
	// for{
	// 	fmt.Println("infinite loop")
	// }

	for j := 0; j < 5; j++ {
		if j == 3 {
			continue
		}
		fmt.Println("j value", j)
	}

	// for range loop
	for k := range 5 {
		fmt.Println("k value", k)
		k++
	}
}
