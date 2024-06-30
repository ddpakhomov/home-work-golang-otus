package main

import (
	"testing"
)

func TestProcessReadings(t *testing.T) {
	c := make(chan int, 10)
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)

	processed := ProcessReadings(c)

	avg := <-processed
	expectedAvg := 4.5 // среднее арифметическое от 0 до 9

	if avg != expectedAvg {
		t.Errorf("expected %v, got %v", expectedAvg, avg)
	}
}
