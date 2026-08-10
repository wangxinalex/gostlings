package main

import "fmt"

func receiveFast(fast, slow <-chan string) string {
	select {
	case value := <-fast:
		return value
	case value := <-slow:
		return value
	}
}

func main() {
	fast := make(chan string, 1)
	slow := make(chan string)
	fast <- "fast lane"
	fmt.Println(receiveFast(fast, slow))
}
