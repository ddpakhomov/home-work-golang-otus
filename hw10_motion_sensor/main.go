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
		max := big.NewInt(100)
		timer := time.After(time.Minute)
		for {
			select {
			case <-timer:
				close(c)
				return
			case c <- func() int {
				if randNum, err := rand.Int(rand.Reader, max); err == nil {
					return int(randNum.Int64())
				}
				return 0
			}():
				time.Sleep(time.Millisecond * 600)
			}
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
		close(processed)
	}()
	return processed
}

func main() {
	readings := SensorReadings()
	processed := ProcessReadings(readings)

	for p := range processed {
		fmt.Println(p)
	}
}
