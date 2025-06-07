package main

import (
	"fmt"
	"time"
)

// worker has to receive jobs channel to read existing jobs
// worker has to receive results channel to send results from jobs
func worker(workerId int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker:", workerId, "starting job:", j)
		time.Sleep(time.Second)
		fmt.Println("worker:", workerId, "finishing job:", j)
		results <- j * 2
	}
}

func main() {
	numJobs := 3
	numWorkers := 2

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// spinned up 3 go routines/workers ready to read from jobs channel
	for w := 0; w < numWorkers; w++ {
		go worker(w, jobs, results)
	}

	for j := 0; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// ensures that goroutines have finished
	// collect resutls
	for range numJobs {
		<-results
	}

}
