package main

import "fmt"

func generate(values ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, value := range values {
			out <- value
		}
	}()
	return out
}

func main() {
	for value := range generate(1, 2, 3) {
		fmt.Println(value)
	}
}
