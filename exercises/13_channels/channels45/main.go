// Advanced pattern: caller-supplied tokens can pace a stream without hiding the
// timing source inside the function.
// Problem: work must proceed only when an external budget or rate permit arrives.
// Without this pattern: the producer runs as fast as possible or owns a timer it
// cannot coordinate with the caller.
// Channels: in carries data; tokens carries one permit per allowed value; out is
// owned and closed by the forwarding goroutine.
// Timeline:
//   receive input
//   receive one token
//   send output
//   repeat
//   close(out)
//
// Hint:
//   Receive one value from in.
//   Receive one token from tokens.
//   Send that value to out.
//   The output owner defers close(out).
//   The caller supplies tokens instead of the function constructing them.

package main

import "fmt"

func rateLimit(tokens <-chan struct{}, in <-chan int) <-chan int {
	// TODO: consume one caller-supplied token per forwarded input and close out
	// after input closes.
	return nil
}

func main() { fmt.Println(rateLimit(nil, nil)) }
