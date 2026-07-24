// Concept: select with a default case — non-blocking channel operations
// Task: the channel is empty (no goroutine ever sends); add a default case so we don't block forever
// Expected output: no value available, doing other work
// Hint: select { case v := <-ch: ... default: ... } — the default branch runs immediately when no channel is ready (Go Tour: Concurrency 5-6)

package main

import "fmt"

func main() {
	ch := make(chan int)

	select {
	case val := <-ch:
		fmt.Println("received:", val)
	default:
		fmt.Println("no value available, doing other work")
	}
}
