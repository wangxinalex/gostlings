package main

import "fmt"

func drain(first, second <-chan int) []int {
	var values []int
	for first != nil || second != nil {
		select {
		case value, ok := <-first:
			if !ok {
				first = nil
				continue
			}
			values = append(values, value)
		case value, ok := <-second:
			if !ok {
				second = nil
				continue
			}
			values = append(values, value)
		}
	}
	return values
}

func main() {
	first := make(chan int, 2)
	second := make(chan int, 2)
	first <- 1
	first <- 3
	second <- 2
	second <- 4
	close(first)
	close(second)
	fmt.Println(drain(first, second))
}
