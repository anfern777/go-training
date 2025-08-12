package main

import (
	"fmt"
	"sync"
)

func squareNumbersWorker(num int) int {
	return num * num
}

func main() {
	w := 4
	j := 100
	var wg sync.WaitGroup
	inputCH := make(chan int)
	outputCH := make(chan int, j)

	wg.Add(w)
	for range w {
		go func() {
			for num := range inputCH {
				outputCH <- squareNumbersWorker(num)
			}
			defer wg.Done()
		}()
	}

	for i := 1; i <= j; i++ {
		inputCH <- i
	}
	close(inputCH)
	wg.Wait()
	close(outputCH)
	sum := 0
	for r := range outputCH {
		sum += r
	}
	fmt.Println(sum)
}
