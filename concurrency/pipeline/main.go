package main

import "fmt"

func main() {
	iterations := 10
	inputs := make(chan int, iterations)

	for i := 0; i < iterations; i++ {
		inputs <- i
	}
	close(inputs)

	outputChannel := NumericPipeline(inputs, func(input int) int { return input * 2 }, func(input int) int { return input * input })

	for i := range outputChannel {
		fmt.Println(i)
	}
}

type NumericOperation func(input int) int

func NumericPipeline(inputs chan int, ops ...NumericOperation) chan int {
	var pipeChannels []chan int
	pipeChannels = append(pipeChannels, inputs)
	for i, op := range ops {
		pipeChannels = append(pipeChannels, make(chan int))
		go func() {
			defer close(pipeChannels[i+1])
			for input := range pipeChannels[i] {
				pipeChannels[i+1] <- op(input)
			}
		}()
	}
	return pipeChannels[len(pipeChannels)-1]
}
