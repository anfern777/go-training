package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	numIncs := 100
	var wg sync.WaitGroup

	container := &Container{
		counters: map[string]int{
			"a": 0,
			"b": 0,
		},
	}

	doIncrement := func(name string, incNum int) {
		for range incNum {
			container.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", numIncs)
	go doIncrement("a", numIncs)
	go doIncrement("b", numIncs)

	wg.Wait()

	fmt.Println(container.counters)

}
