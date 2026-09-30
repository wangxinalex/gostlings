// Concept: context.WithCancel stops cooperative work.
// Task: return the cancellation result as soon as ctx.Done() is closed.
// Hint: wait for the context's cancellation channel and the work gate in one
//       select; a canceled context must win without waiting for work to be released.
// Stuck?: ctx.Done is the cancellation channel; return the canceled result from that case.

package main

import (
	"context"
	"fmt"
)

var workGate = make(chan struct{})

func worker(ctx context.Context) string {
	// TODO: Select on ctx.Done() so cancellation stops the worker.
	return "worker: completed"
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	fmt.Println(worker(ctx))
}
