// Concept: closing a channel tells range loops to stop
// Task: this program deadlocks because range never stops; close the channel after sending
// Expected output: 1
// 2
// 3
// Hint: close(ch) after the sends so the range loop knows there are no more values (Go Tour: Concurrency 4)

package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		// 思路：range 需要知道“不会再有值”。只有发送方发送完最后一个值后关闭，
		// 接收方才会先读完已有值，再让 range 正常结束。
		// TODO: Close the channel so the range below does not deadlock.
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
