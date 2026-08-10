// Concept: buffered channels decouple send and receive
// Task: this program deadlocks because the channel is unbuffered; change only the make line to fix it
// Expected output: 1
// 2
// Hint: make(chan int, 2) creates a buffered channel that holds 2 values without blocking (Go Tour: Concurrency 3)

package main

import "fmt"

func main() {
	// 思路：容量为 2 的缓冲区可以暂存两个值，让发送先完成；容量耗尽后仍然会阻塞。
	ch := make(chan int) // TODO: Make this buffered so the sends don't block.

	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
