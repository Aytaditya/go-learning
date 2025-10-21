package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	// inside the defer function which means it will be executed at the end of the function
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock() // lock the critical section (shared resources)
	p.views += 1
}

func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)

	}

	wg.Wait()
	fmt.Println(myPost.views)
}
