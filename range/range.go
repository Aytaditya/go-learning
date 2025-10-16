package main

import "fmt"

func main() {
	var nums = []int{2, 3, 4}
	sum := 0
	// _ is ehere we get index
	for _, num := range nums {
		sum = sum + num
	}
	fmt.Println("sum:", sum)

	m := map[int]int{1: 10, 2: 20, 3: 30}
	for _, value := range m {
		println(value)
	}
}
