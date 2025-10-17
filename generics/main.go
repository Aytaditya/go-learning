package main

// the problem is this we can only print slice of int but what if we want to print slice of string or float64
func print(items []int) {
	for _, v := range items {
		println(v)
	}
}

// we can use generics to solve this problem
func genericPrint[T any](items []T) {
	for _, v := range items {
		println(v)
	}
}

// now instead of making we can make it custom for example int and float32
func customPrint[T int | float32](items []T) {
	for _, v := range items {
		println(v)
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	names := []string{"Alice", "Bob", "Charlie"}
	flNums := []float32{1.1, 2.2, 3.3}
	print(nums)
	genericPrint(nums)
	genericPrint(names)

	customPrint(nums)
	customPrint(flNums)
	// customPrint(names)  this wont work as names is string type
}
