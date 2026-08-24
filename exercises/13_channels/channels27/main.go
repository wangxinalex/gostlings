// Problem: parallel workers finish in nondeterministic order, but an API may
// promise output order matching the input slice.
// Without this pattern: appending results as they arrive changes the caller's
// ordering contract.
// Channels: indexed jobs and indexed results carry data; the results closer is
// still coordinator-owned. The index is metadata, not a synchronization signal.
// Timeline: input index -> worker completes in any order -> collector stores by index -> return input order
// Hint: put the original index in each job and result. The collector writes
// result.value into out[result.index] instead of appending in receive order.
// Reuse the basic pool's producer, worker acknowledgements, and results close.

package main

import "fmt"

var processOrderedJob = func(value int) int { return value * value }

func runOrdered(workers int, jobs []int) []int {
	return nil // TODO: carry indexes through the pool and restore input order
}

func main() {
	fmt.Println(runOrdered(2, []int{4, 1, 3, 2}))
}
