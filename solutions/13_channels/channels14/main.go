package main

import "fmt"

func runAsync(work func() int) <-chan int {
	out := make(chan int, 1)
	go func() {
		out <- work()
		close(out)
	}()
	return out
}

func main() {
	fmt.Println(<-runAsync(func() int { return 42 }))
}
