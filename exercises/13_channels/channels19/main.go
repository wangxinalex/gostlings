// Concept: a pipeline stage owns and closes its output
// Task: square every input value and close the output after input closes
// Expected behavior: the stage can be ranged over until completion
// Hint: one goroutine ranges in, sends transformed values, and defers close(out)

package main

import "fmt"

func square(in <-chan int) <-chan int {
	// 思路：每个 stage 只关闭自己的输出；它从不关闭调用方提供的输入。
	return nil // TODO: transform values in a goroutine and close the output
}

func main() {
	in := make(chan int, 3)
	in <- 1
	in <- 2
	in <- 3
	close(in)
	for value := range square(in) {
		fmt.Println(value)
	}
}
