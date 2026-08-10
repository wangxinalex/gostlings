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
	fmt.Println(await(make(chan string)))
}
