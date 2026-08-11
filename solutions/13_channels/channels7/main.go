package main

import "fmt"

func receiveAll(ch <-chan int) []int {
	var values []int
	for value := range ch {
		values = append(values, value)
	}
	return values
}

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
	for _, value := range receiveAll(generate(1, 2, 3)) {
		fmt.Println(value)
	}
}
