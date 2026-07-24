// Concept: returning and checking errors
// Task: make the divide function return an error when b is 0, and check that error in main
// Expected output: divisor must not be 0
// Hint: use errors.New() to create an error value (Go Tour: Methods 19)

package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divisor must not be 0")
	}
	return a / b, nil
}

func main() {
	_, err := divide(6, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("no error")
}
