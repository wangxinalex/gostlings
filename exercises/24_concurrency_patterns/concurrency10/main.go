// Concept: fan-in needs one coordinator to close a shared output.
// Task: merge typed result streams while preserving failures and cancellation.
// Hint: one forwarder per source, an acknowledgment per forwarder, one closer.
package main

import "context"

type mergeResult struct {
	value int
	err   error
}

func merge(ctx context.Context, sources ...<-chan mergeResult) <-chan mergeResult {
	// TODO: Forward all sources and close the output after every source stops.
	return nil
}
