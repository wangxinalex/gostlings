// Concept: select multiplexes channel operations
// Task: c1 sends a value after a short delay and c2 never does; use select to receive from c1 before the timeout
// Expected output: received: fast lane
// Hint: a select statement picks whichever case is ready first (Go Tour: Concurrency 5)

package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(10 * time.Millisecond)
		c1 <- "fast lane"
	}()

	// c2 never receives a value — only c1 is ready.

	// TODO: Use select to receive from whichever channel is ready first.
	fmt.Println(<-c1)
}
