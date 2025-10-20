package main

import (
	"fmt"
	"time"
)

// channels are used to communicate between goroutines

func process(numChannel chan int) {
	fmt.Println("Processing numbers...", <-numChannel)
}

func sum(result chan int, num1 int, num2 int) {
	ans := num1 + num2
	result <- ans // send the result to the channel
}

// emailChannel <- chan	(this means only receive from channel no sending is allowed)
// chan <- emailChannel (this means only sending to channel no receiving is allowed)
func EmailSender(emailChannel <-chan string, done chan<- bool) {
	// function to send email
	defer func() { done <- true }() // this will be run at the last as cleaning function to notify that the goroutine is done
	// emailChannel<-"hii" this will cause compile time error because the channel is receive only
	for email := range emailChannel {
		fmt.Println("Sending email to:", email)
	}
}

func main() {
	// messageChannel := make(chan string)  // create a channel of type string
	// messageChannel <- "Hello, Channels!" // send a message to the channel

	// msg := <-messageChannel // receive a message from the channel
	// fmt.Println(msg)

	// now lets make a buffered channel (better for queue system)
	emailChannel := make(chan string, 100) // buffered channel with capacity of 100

	done := make(chan bool) // channel as a alternative to wait group

	// emailChannel <- "firstEmail@google.com"
	// emailChannel <- "twoEmail@google.com"

	// this will not block because the channel is buffered (no deadlock will occur)
	// fmt.Println(<-emailChannel)
	// fmt.Println(<-emailChannel)

	go EmailSender(emailChannel, done) // start the email sender goroutine

	for i := 0; i < 10; i++ {
		emailChannel <- fmt.Sprintf("%d@gmail.com", i)
	}
	close(emailChannel) // close the channel after sending all emails (prevents deadlock in the goroutine)
	<-done              // wait for the goroutine to finish

	// this is a unbuffered channel example (only one value can be sent at a time)
	numChannel := make(chan int)
	go process(numChannel)
	numChannel <- 42

	result := make(chan int)
	go sum(result, 10, 20)

	res := <-result // receive the result from the channel
	fmt.Println("Sum is:", res)

	// working on multiple channels
	chan1 := make(chan int)
	chan2 := make(chan string)

	// inline goroutine
	go func() {
		chan1 <- 100
	}() // last bracket to call the function

	go func() {
		chan2 <- "Hello from chan2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case channel1Val := <-chan1:
			fmt.Println("Value from chan1:", channel1Val)
		case channel2Val := <-chan2:
			fmt.Println("Value from chan2:", channel2Val)
		}
	}

	time.Sleep(1 * time.Second) // wait for goroutine to finish

}
