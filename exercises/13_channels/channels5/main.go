// Concept: generator — the producer owns and closes its output channel
// Task: send every input value from a goroutine, then close the returned channel
// Expected behavior: callers can range over the result, including for empty input
// Hint: defer close(out) inside the producer goroutine; the caller only receives

package main

import "fmt"

func generate(values ...int) <-chan int {
	// 思路：返回只读 channel，把发送和关闭责任留在生产者内部；
	// 调用方只需要 range，不需要猜测何时关闭，也不能误关闭它。
	return nil // TODO: create the output, send values in a goroutine, and close it
}

func main() {
	for value := range generate(1, 2, 3) {
		fmt.Println(value)
	}
}
