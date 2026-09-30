// Concept: splitting a slice across per-goroutine partial sums
// Task: sum every value with at most workers goroutines and add the partial totals after joining
// Expected behavior: even and uneven splits and more workers than values all return the arithmetic
//                  total, and an empty input returns 0.
// Hint: give every worker its own slot in a slice of partial totals, so each goroutine writes only
//       its own slot, owns its partial total, and no mutex is needed. Clamp the worker count to the
//       input length, join every goroutine with sync.WaitGroup, and add the partial totals only
//       after the join returns.
// Version note: this repository targets Go 1.26, so sync.WaitGroup.Go is also available;
// older Go versions use Add, go, and defer Done.

package main

func sumParallel(values []int, workers int) int {
	// TODO: Partition values across at most workers goroutines, let each goroutine sum
	// its own partition into its own partial total, join them, and return the sum of
	// the partial totals.
	return 0
}

func main() {}
