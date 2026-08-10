// Concept: nil channels disable select cases
// Task: drain two inputs and set each one to nil after it closes
// Expected behavior: all values arrive once, with no repeated zero values after close
// Hint: in a select loop, assign a closed input variable to nil to remove that case

package main

import "fmt"

func drain(first, second <-chan int) []int {
	// 思路：closed channel 永远可读，会不断产生零值；把对应变量设为 nil，
	// select 就会永久禁用这个 case，直到另一个输入也结束。
	return nil // TODO: disable each input after its close and collect values
}

func main() {
	first := make(chan int, 2)
	second := make(chan int, 2)
	first <- 1
	first <- 3
	second <- 2
	second <- 4
	close(first)
	close(second)
	fmt.Println(drain(first, second))
}
