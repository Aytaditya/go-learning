package main

// this function can recieve n number of integers as arguments (add is now a variadic function)
func add(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	result := add(1, 2, 3, 4, 5)
	println(result) // 15
}
