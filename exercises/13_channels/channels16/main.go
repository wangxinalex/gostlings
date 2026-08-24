// Problem: a producer may keep generating after the consumer has abandoned its
// output, leaving the producer blocked on the next send.
// Without this pattern: closing stop does not interrupt a plain `out <- value`.
// Channels: stop is a receive-only cancellation signal; out carries values and
// is closed by the producer when it exits.
// Timeline: produce -> select send or stop -> close(out) on normal finish/cancel
// Hint: put every output send in a select with `<-stop`. Defer close(out) inside
// the producer goroutine so both cancellation and normal completion close it.

package main

import "fmt"

func produce(stop <-chan struct{}) <-chan int {
	return nil // TODO: send values with a stop case and close out when stopping
}

func main() {
	stop := make(chan struct{})
	out := produce(stop)
	fmt.Println(<-out)
	close(stop)
}
