package main

import (
	"fmt"
	"maps"
)

func main() {
	// similar to hash map, objects , and dictionary
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	println(m["a"])
	delete(m, "a")
	clear(m) // clears the map

	mapp := map[string]int{"one": 1, "two": 2}
	mapp1 := map[string]int{"one": 1, "two": 2}
	//println(mapp["one"])

	v, ok := mapp["one"] // The variable ok returned by a map lookup is a boolean, not a number and value will be stored in v variable
	if ok && v > 0 {
		fmt.Println("exists", v)
	}

	fmt.Println(maps.Equal(mapp, mapp1)) // true
}
