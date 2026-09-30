package main

import "fmt"

func drain(first, second <-chan int) []int {
	values := []int{}
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

func main() { fmt.Println(drain(nil, nil)) }
