package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(workerId int, jobs <-chan int) {
	for j := range jobs {
		fmt.Println("worker:", workerId, "starting job:", j)
		time.Sleep(time.Second)
		fmt.Println("worker:", workerId, "finishing job:", j)
	}
}

func main() {
	var wg sync.WaitGroup
	const numJobs = 10
	const numWorkers = 3
	jobs := make(chan int, numJobs)

	// spin up workers, i.e. wait groups
	for w := range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(w, jobs)
		}()
	}

	for j := range numJobs {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
}
