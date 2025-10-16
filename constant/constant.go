package main

import "fmt"

func main() {
	const name = "Golang" //once assigned cannot be changed
	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(name)
	fmt.Println(port)

}
