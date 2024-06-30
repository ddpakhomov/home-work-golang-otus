package main

import (
	"fmt"
	"math/rand"
	"time"
)

func SensorReadings() <-chan int {
	c := make(chan int)
	go func() {
		for {
			c <- rand.Intn(100)
			time.Sleep(time.Millisecond * 600)
		}
	}()
	return c
}

func ProcessReadings(c <-chan int) <-chan float64 {
	processed := make(chan float64)
	go func() {
		var sum int
		var count int
		for reading := range c {
			sum += reading
			count++
			if count == 10 {
				processed <- float64(sum) / 10
				sum = 0
				count = 0
			}
		}
	}()
	return processed
}

func main() {
	readings := SensorReadings()
	processed := ProcessReadings(readings)

	startTime := time.Now()
	iterations := 0

	for time.Since(startTime) < time.Minute {
		fmt.Println(<-processed)
		iterations++
	}
	fmt.Println("Минута прошла. Количество выполненных итераций:", iterations)
}
