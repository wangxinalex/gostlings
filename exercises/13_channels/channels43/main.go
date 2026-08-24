// Advanced pattern: a buffered token channel acts as a semaphore-like limit.
// Problem: a pool must cap active work even when the input contains many jobs.
// Without this pattern: all jobs may call work concurrently and overload a
// dependency or local resource.
// Channels: tokens carries capacity permits, not business data; each worker
// receives one token before work and sends it back afterward. Indexed results
// restore input order.
// Timeline: acquire token -> work -> release token -> store indexed result
// Hint: prefill a buffered tokens channel with limit empty structs. Receive one
// token before work and return it afterward. Carry each job's index with its
// result and store results by index.
package main

import "fmt"

func parallel(limit int, jobs []int, work func(int) int) []int {
	return nil // TODO: use a buffered token channel to bound work and index results to restore order
}

func main() { fmt.Println(parallel(1, nil, func(value int) int { return value })) }
