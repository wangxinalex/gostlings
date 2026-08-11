// Concept: a buffered channel can hold a fixed number of semaphore tokens.
// Task: run work for every job without exceeding limit active calls, then restore input order.
// Expected behavior: at most limit calls to work run at once and returned values match job order.
// Hint: prefill a buffered tokens channel with limit empty structs. A goroutine receives a token before work
// and returns it afterward; carry each job's index with its result and store results by index.
package main

import "fmt"

func parallel(limit int, jobs []int, work func(int) int) []int {
	return nil // TODO: use a buffered token channel to bound work and index results to restore order
}

func main() { fmt.Println(parallel(1, nil, func(value int) int { return value })) }
