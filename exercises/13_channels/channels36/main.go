// Problem: a pipeline must transport either a successful value or an error
// without losing which outcome belongs to a result item.
// Without this pattern: separate value and error channels can arrive out of
// order, forcing callers to reconstruct pairs.
// Channels: each input and out carries a complete result envelope; exited is a
// lifecycle acknowledgement and never carries application data.
// Timeline: forward value+error unchanged -> forwarder exits -> coordinator counts -> close(out)
// Hint: start one forwarder per input and send the complete result value without
// changing either field. Use one raw exit acknowledgement per forwarder so one
// coordinator closes out after all inputs finish.
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
