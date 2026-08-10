// Concept: cancellable fan-in — downstream abandonment must stop forwarders
// Task: merge inputs, but stop receiving and sending when stop is closed
// Expected behavior: a blocked input cannot keep the merged output alive after cancellation
// Hint: select on stop while receiving from each input and while sending to out

package main

import "fmt"

func merge(stop <-chan struct{}, inputs ...<-chan int) <-chan int {
	// 思路：只在 out <- value 上响应 stop 还不够；forwarder 也可能先阻塞在 <-input，
	// 所以接收和发送两侧都必须可取消。
	return nil // TODO: add cancellable forwarders and close out after they exit
}

func main() {
	input := make(chan int, 2)
	input <- 1
	input <- 2
	close(input)
	for value := range merge(make(chan struct{}), input) {
		fmt.Println(value)
	}
}
