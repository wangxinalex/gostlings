// Concept: fan-in needs one coordinator to close a shared output.
// Task: merge typed result streams while preserving failures and cancellation.
// Hint: one forwarder per source, an acknowledgment per forwarder, one closer.
package main

import (
	"context"
	"sync"
)

type mergeResult struct {
	value int
	err   error
}

func merge(ctx context.Context, sources ...<-chan mergeResult) <-chan mergeResult {
	out := make(chan mergeResult)
	var wg sync.WaitGroup
	wg.Add(len(sources))
	for _, source := range sources {
		go func(source <-chan mergeResult) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case item, ok := <-source:
					if !ok {
						return
					}
					select {
					case out <- item:
					case <-ctx.Done():
						return
					}
				}
			}
		}(source)
	}
	go func() { wg.Wait(); close(out) }()
	return out
}

func main() {}
