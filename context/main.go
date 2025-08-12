package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Response struct {
	num int
	err error
}

func myThridPartyRequestThatCanBeHeavy() (int, error) {
	time.Sleep(200 * time.Millisecond)
	return 100, nil
}

func fetchUserData(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 1000*time.Millisecond)
	defer cancel()

	resCH := make(chan Response)
	go func() {
		res, err := myThridPartyRequestThatCanBeHeavy()
		resCH <- Response{
			num: res,
			err: err,
		}
	}()

	for {
		select {
		// moment that context timeout is triggered
		case <-ctx.Done():
			return -1, fmt.Errorf("fetching data from third party took too long")
		case resp := <-resCH:
			return resp.num, nil
		}
	}
}

func main() {
	start := time.Now()
	res, err := fetchUserData(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	println("result: ", res)
	fmt.Println("took: ", time.Since(start))
}
