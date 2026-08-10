// Concept: multi-stage pipelines compose independent channel lifecycles
// Task: build a pipeline that doubles each value and then adds one
// Expected behavior: [1,2,3] becomes [3,5,7], then output closes
// Hint: implement transform once, then compose two transform stages

package main

import "fmt"

func transform(in <-chan int, fn func(int) int) <-chan int {
	return nil // TODO: range in, apply fn, send out, and close out
}

func pipeline(in <-chan int) <-chan int {
	// 思路：上游关闭自己的 output 后，下游的 range 才能结束；
	// 每个 stage 都要独立负责“读到结束、关闭输出”的协议。
	return nil // TODO: compose double and add-one stages
}

func main() {
	in := make(chan int, 3)
	in <- 1
	in <- 2
	in <- 3
	close(in)
	for value := range pipeline(in) {
		fmt.Println(value)
	}
}
