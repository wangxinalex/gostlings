// Concept: worker completion order differs from input order, so results carry their index.
// Task: square jobs in parallel and return results in the same order as jobs.
// Expected behavior: values may finish in any order, but runOrdered restores each value to its input position.
// Hint: send indexed jobs to workers and send an indexed result back. The collector writes result.value into out[result.index].
//       Keep the same producer, worker acknowledgements, and coordinator-owned results close as the basic pool.

package main

import "fmt"

var processOrderedJob = func(value int) int { return value * value }

func runOrdered(workers int, jobs []int) []int {
	return nil // TODO: carry indexes through the pool and restore input order
}

func main() {
	fmt.Println(runOrdered(2, []int{4, 1, 3, 2}))
}
