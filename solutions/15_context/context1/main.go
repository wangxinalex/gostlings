// Concept: context.WithCancel stops cooperative work.
// Task: return the cancellation result as soon as ctx.Done() is closed.
// Hint: select on ctx.Done() and the work gate; a canceled context must win
// without waiting for work to be released.
package main

import (
	"context"
	"fmt"
)

var workGate = make(chan struct{})

func worker(ctx context.Context) string {
	select {
	case <-ctx.Done():
		return "worker: canceled"
	case <-workGate:
		return "worker: completed"
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cancel()
	fmt.Println(worker(ctx))
}
