package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func SensorReadings() <-chan int {
	c := make(chan int)
	go func() {
		for {
			max := big.NewInt(100)
			randNum, _ := rand.Int(rand.Reader, max)
			c <- int(randNum.Int64())
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
