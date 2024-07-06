package main

import (
	"fmt"
	"sync"
)

var (
	counter = 0
	lock    = sync.Mutex{}
)

func Worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	lock.Lock()
	counter++
	fmt.Printf("Worker %d has finished task. Counter: %d\n", id, counter)
	lock.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Worker(&wg, i)
	}

	wg.Wait()
	fmt.Println("All workers have finished their tasks.")
}
