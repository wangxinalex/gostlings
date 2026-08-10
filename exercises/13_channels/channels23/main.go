// Concept: ticker channels can act as a rate limiter
// Task: forward one input only after one tick is received
// Expected behavior: every input consumes one tick, then output closes
// Hint: wait on <-ticks before sending each value; the caller owns the ticker lifecycle

package main

import (
	"fmt"
	"time"
)

func rateLimit(ticks <-chan time.Time, in <-chan int) <-chan int {
	// 思路：把“允许开始下一次工作”的事件建模成 channel，
	// 生产速度就受 ticker 节奏约束，而不是靠固定 Sleep。
	return nil // TODO: wait for a tick before forwarding each input
}

func main() {
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()
	in := make(chan int, 2)
	in <- 1
	in <- 2
	close(in)
	for value := range rateLimit(ticker.C, in) {
		fmt.Println(value)
	}
}
