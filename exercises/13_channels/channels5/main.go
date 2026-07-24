// Concept: select with a time.After timeout — don't block forever
// Task: the select blocks on a slow channel that never fires; add a time.After timeout case
// Expected output: timed out
// Hint: case <-time.After(100 * time.Millisecond): fires after that duration and sets an upper bound on wait (Go Tour: Concurrency 5-6)

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	// This goroutine takes way too long.
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "late result"
	}()

	// TODO: Add a time.After case to the select so it prints "timed out" after 100ms.
	select {
	case msg := <-ch:
		fmt.Println(msg)
	}
}
