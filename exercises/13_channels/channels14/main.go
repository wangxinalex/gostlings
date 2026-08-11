// Concept: asynchronous result — a one-shot buffered result lets a worker finish first
// Task: run work in a goroutine and return its result channel immediately
// Expected behavior: work can publish its one result before the caller receives it
// Hint: make out with capacity 1, then start a goroutine that sends work() to out.
//       A capacity of 1 prevents the worker from blocking just because the caller is late.

package main

import "fmt"

func runAsync(work func() int) <-chan int {
	return nil // TODO: return a capacity-one result channel and run work in a goroutine
}

func main() {
	fmt.Println(<-runAsync(func() int { return 42 }))
}
