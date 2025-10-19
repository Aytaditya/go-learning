package main

import (
	"fmt"
	"time"
)

// channels are used to communicate between goroutines

func process(numChannel chan int) {
	fmt.Println("Processing numbers...", <-numChannel)
}

func main() {
	// messageChannel := make(chan string)  // create a channel of type string
	// messageChannel <- "Hello, Channels!" // send a message to the channel

	// msg := <-messageChannel // receive a message from the channel
	// fmt.Println(msg)

	numChannel := make(chan int)
	go process(numChannel)
	numChannel <- 42

	time.Sleep(1 * time.Second) // wait for goroutine to finish

}
