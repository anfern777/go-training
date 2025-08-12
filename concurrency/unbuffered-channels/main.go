package main

import (
	"fmt"
	"time"
)

// func main() {
// 	ch := make(chan int)

// 	// set up go routine to be ready to read from ch
// 	go func() {
// 		for iteration := range ch {
// 			done := make(chan bool)
// 			go executeMethod(iteration, done)
// 			<-done
// 		}
// 	}()
// 	for i := 0; i < 3; i++ {
// 		ch <- i
// 	}
// }

// func executeMethod(cur int, done chan bool) {
// 	for _, a := range []string{"a", "b", "c"} {
// 		fmt.Printf("current iteration: %d; current msg: %s\n", cur, a)
// 		done <- true
// 	}
// }

func main() {
	done := make(chan bool)
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Hello from a goroutine!") // this won't ever be reached
		done <- true
	}()
	time.Sleep(2 * time.Second)
	<-done
	fmt.Println("I am also a goroutine, but a special one!")
}
