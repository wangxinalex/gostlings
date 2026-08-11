// Concept: a caller can supply channel tokens to rate-limit work without constructing a ticker.
// Task: forward one input value for every received token.
// Expected behavior: values wait for tokens and output closes after input closes and its final value is forwarded.
// Hint: receive one value from in, then receive one token from tokens before sending that value to out. The
// goroutine that owns out defers close(out); tokens are supplied by the caller rather than a constructed clock.
package main

import "fmt"

func rateLimit(tokens <-chan struct{}, in <-chan int) <-chan int {
	return nil // TODO: consume one caller-supplied token per forwarded input and close out after input closes
}

func main() { fmt.Println(rateLimit(nil, nil)) }
