package main

import (
	"fmt"
	"time"
)

func main() {
	var result, i uint64
	// single goroutine executes calc
	start := time.Now()
	for i = 1; i < 10000000000; i++ {
		result += i
	}
	elapsed := time.Since(start)
	fmt.Print("Consumed time：", elapsed)
	fmt.Print(", result：", result, "\n")

	// four goroutines execute calc
	start = time.Now()
	ch1 := calc(1, 2500000000)
	ch2 := calc(2500000001, 5000000000)
	ch3 := calc(5000000001, 7500000000)
	ch4 := calc(7500000001, 10000000000)
	result = <-ch1 + <-ch2 + <-ch3 + <-ch4
	elapsed = time.Since(start)
	fmt.Print("Consumed time：", elapsed)
	fmt.Print(", result：", result)

}

func calc(from uint64, to uint64) <-chan uint64 {
	ch := make(chan uint64)
	go func() {
		result := from
		for i := from + 1; i <= to; i++ {
			result += i
		}
		ch <- result
	}()
	return ch
}
