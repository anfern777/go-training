package main

import "fmt"

func main() {
	ch := make(chan int)

	ch <- 3
	ch <- 1

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
