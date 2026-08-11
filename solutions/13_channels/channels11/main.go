package main

import (
	"fmt"
	"time"
)

func await(ch <-chan string) string {
	select {
	case value := <-ch:
		return value
	case <-time.After(50 * time.Millisecond):
		return "timed out"
	}
}

func main() {
	result := make(chan string, 1)
	result <- "ready"
	fmt.Println(await(result))
}
