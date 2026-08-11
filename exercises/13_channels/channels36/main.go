// Concept: result envelopes keep successful values and errors on the same stream.
// Task: merge result streams without changing either field of any envelope.
// Expected behavior: values and errors both arrive unchanged; empty input closes output.
// Hint: start one forwarder per input, have each send its complete result value, then use raw exit
//       acknowledgements so one coordinator closes out after every forwarder has returned.
package main

import "fmt"

type result struct {
	value int
	err   error
}

func mergeResults(inputs ...<-chan result) <-chan result {
	return nil // TODO: forward envelopes and coordinate output close
}
func main() { fmt.Println(mergeResults()) }
