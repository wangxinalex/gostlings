package main

import "fmt"

func watch(done <-chan struct{}) <-chan string {
	out := make(chan string, 1)
	go func() { <-done; out <- "done"; close(out) }()
	return out
}
func main() { fmt.Println(watch(make(chan struct{}))) }
