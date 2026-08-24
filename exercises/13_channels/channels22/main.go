// Problem: several independent producers need to look like one stream to a
// downstream consumer.
// Without this pattern: the consumer must know every input and cannot range one
// output until all producers finish.
// Channels: each input is receive-only and producer-owned; out is created and
// closed by merge; exited carries one acknowledgement per forwarder.
// Timeline: inputs -> one forwarder each -> shared out -> coordinator closes out after all exit tokens
// Hint: start one forwarder goroutine per input. It ranges its input and sends
// values to out, then sends one buffered acknowledgement. A separate coordinator
// receives len(inputs) acknowledgements and is the only goroutine that closes out.
// With no inputs, return an already-closed output.

package main

import "fmt"

var onForwarderExit = func() {}

func merge(inputs ...<-chan int) <-chan int {
	return nil // TODO: forward every input and close out from one coordinator
}

func main() {
	first := make(chan int, 1)
	second := make(chan int, 1)
	first <- 1
	second <- 2
	close(first)
	close(second)
	for value := range merge(first, second) {
		fmt.Println(value)
	}
}
