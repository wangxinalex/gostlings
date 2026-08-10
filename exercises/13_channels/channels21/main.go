// Concept: cancellation-aware pipelines stop both receives and sends
// Task: apply two stages while allowing stop to interrupt a blocked input or output
// Expected behavior: closing stop closes the final output and lets every stage exit
// Hint: select on stop when receiving from input and sending to output in every stage

package main

import "fmt"

func pipeline(stop <-chan struct{}, in <-chan int) <-chan int {
	// 思路：pipeline 是多个 goroutine 的链条；只取消最后一层会让上游继续卡住，
	// 所以同一个 stop 必须传到每个 stage 的接收和发送位置。
	return nil // TODO: compose cancellable stages
}

func main() {
	in := make(chan int, 2)
	in <- 1
	in <- 2
	close(in)
	for value := range pipeline(make(chan struct{}), in) {
		fmt.Println(value)
	}
}
