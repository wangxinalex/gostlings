// Problem: one server handles many requests, but every caller needs its own
// response route.
// Without this pattern: a shared result stream forces callers to coordinate
// which response belongs to which request.
// Channels: requests carries request data and a private reply channel; reply is
// owned by the request sender as the destination; done is closed by the server
// after requests closes.
// Timeline:
//   caller sends request and reply channel
//   server computes
//   server sends on the private reply channel
//   caller closes requests
//   server closes done
//
// Hint:
//   Range over requests in one server goroutine.
//   Send value*2 on each request.reply.
//   Defer close(done).
//   The reply channel routes data; it is not a global shutdown signal.

package main

import "fmt"

type request struct {
	value int
	reply chan int
}

func serve(requests <-chan request) <-chan struct{} {
	return nil // TODO: reply to every request and close done after input closure
}

func main() { fmt.Println(serve(nil)) }
