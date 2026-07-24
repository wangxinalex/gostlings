// Concept: buffered channels decouple send and receive
// Task: this program deadlocks because the channel is unbuffered; change only the make line to fix it
// Expected output: 1
// 2
// Hint: make(chan int, 2) creates a buffered channel that holds 2 values without blocking (Go Tour: Concurrency 3)

package main

import "fmt"

func main() {
	ch := make(chan int) // TODO: Make this buffered so the sends don't block.

	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
