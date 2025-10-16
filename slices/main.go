package main

import (
	"fmt"
	"sort"
)

func main() {
	// declared a slice of integers (1 way to make slice)
	var nums []int
	for i := 0; i < 5; i++ {
		nums = append(nums, i)
	}
	// to push element in slice we use append function and pass nums and value we want to push
	fmt.Println(nums)
	fmt.Println(len(nums))

	// using make function to create a slice
	slice := make([]int, 5)
	// appending values in slice
	for i := 0; i < 5; i++ {
		slice = append(slice, i)
	}
	// setting values in slice
	for i := 0; i < 5; i++ {
		slice[i] = i * 2
	}
	// sorting slice
	sort.Ints(slice)

	fmt.Println(slice)

	// another way to make slice
	hi := []int{}
	fmt.Println(hi)
	fmt.Println(len(hi))
	fmt.Println(cap(hi))

	for i := 0; i < 5; i++ {
		hi = append(hi, i)
	}
	fmt.Println(hi)
	fmt.Println(len(hi))
	fmt.Println(cap(hi))
	fmt.Println(hi[0:2]) // slice operator
}
