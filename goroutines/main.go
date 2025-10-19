package main

import (
	"fmt"
	"sync"
)

// wait group will be passed as pointer
func task(id int, wg *sync.WaitGroup) {
	// defer will be run at the last of the function
	defer wg.Done() // decrementing the wait group counter
	fmt.Println("Task", id, "is running")
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		// if we do wg.Add(2) program wil not exit
		wg.Add(1)       // incrementing the wait group counter
		go task(i, &wg) // now this will run concurrently
	}

	wg.Wait() // waiting for all goroutines to finish

	// sleeping the main goroutine to allow other goroutines to finish
	// time.Sleep(2 * time.Second)
	//we dont use like this bye sleeping instead we use sync.WaitGroup or channels to wait for goroutines to finish
}
