package main

import (
	"sync"
	"testing"
)

func TestWorker(t *testing.T) {
	var wg sync.WaitGroup
	lock.Lock() // Убедитесь, что счетчик сброшен перед запуском теста
	counter = 0
	lock.Unlock()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Worker(&wg, i)
	}
	wg.Wait()

	lock.Lock()
	if counter != 5 {
		t.Errorf("Ожидалось значение счетчика 5, но получили %d", counter)
	}
	lock.Unlock()
}
