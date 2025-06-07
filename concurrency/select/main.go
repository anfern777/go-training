package main

import (
	"fmt"
	"time"
)

// func main() {
// 	c1 := make(chan string)
// 	c2 := make(chan string)

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		c1 <- "one"
// 	}()
// 	go func() {
// 		time.Sleep(10 * time.Second)
// 		c2 <- "two"
// 	}()

// 	for range 2 {
// 		select {
// 		case msg1 := <-c1:
// 			fmt.Println("received", msg1)
// 		case msg2 := <-c2:
// 			fmt.Println("received", msg2)
// 		case <-time.After(5 * time.Second):
// 			fmt.Println("Timeout!")
// 		}
// 	}
// }

func main() {
	msgs := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-msgs:
		fmt.Println("received msg", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hello"
	select {
	case msgs <- msg:
		fmt.Println("sent msg", msg)
	default:
		fmt.Println("no message sent")
	}

	go func() {
		signals <- true
	}()

	time.Sleep(1 * time.Second)

	select {
	case msg := <-msgs:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
		// default:
		// 	fmt.Println("no activity")
	}
}
