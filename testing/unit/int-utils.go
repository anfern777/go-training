package main

import "fmt"

func ReturnGreaterNum(a int, b int) (int, error) {
	if a == b {
		return -1, fmt.Errorf("a and b are equal")
	}
	if a > b {
		return a, nil
	} else {
		return b, nil
	}
}
