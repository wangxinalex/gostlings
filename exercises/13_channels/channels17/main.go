// Concept: parallel workers finish in arbitrary order
// Task: preserve the input order while still processing jobs concurrently
// Expected behavior: runOrdered(4, []int{1,2,3}) returns []int{1,4,9}
// Hint: send each job's index with its result, then write into a pre-sized output slice

package main

import "fmt"

func runOrdered(workers int, jobs []int) []int {
	// Thought: channel receive order is completion order, not business order.
	// Carry the index and place results back by index to restore input order.
	return nil // TODO: add indexes to jobs/results and restore order
}

func main() {
	fmt.Println(runOrdered(2, []int{1, 2, 3, 4}))
}
