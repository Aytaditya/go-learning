package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

// go can return multiple values
func languages(input string) (string, string, string) {
	return "Go", "Python", input
}

// basically we have passed a function as an argument to another function that passed function accespts two integers and returns an integer
func process(fn func(a int, b int) int) int {
	result := fn(3, 4)
	fmt.Println(result)
	return result
}

func main() {
	result := add(1, 2)
	fmt.Println(result)
	fmt.Println(languages("Java"))
	// or we can do this
	// lang1, lang2, lang3 := languages("Java")

	fmt.Println("Passing function as an argument and answer is", process(add))
}
