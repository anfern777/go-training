package main

import "fmt"

func main() {
	ch := make(chan int)

	// set up go routine to be ready to read from ch
	go func() {
		for iteration := range ch {
			done := make(chan bool)
			go executeMethod(iteration, done)
			<-done
		}
	}()
	for i := 0; i < 3; i++ {
		ch <- i
	}
}

func executeMethod(cur int, done chan bool) {
	for _, a := range []string{"a", "b", "c"} {
		fmt.Printf("current iteration: %d; current msg: %s\n", cur, a)
		done <- true
	}
}
