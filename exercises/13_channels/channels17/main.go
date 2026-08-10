// Concept: parallel workers finish in arbitrary order
// Task: preserve the input order while still processing jobs concurrently
// Expected behavior: runOrdered(4, []int{1,2,3}) returns []int{1,4,9}
// Hint: send each job's index with its result, then write into a pre-sized output slice

package main

import "fmt"

func runOrdered(workers int, jobs []int) []int {
	// 思路：channel 收到的顺序是完成顺序，不是业务顺序；把 index 一起传递，
	// 消费结果时按 index 放回去，才能恢复输入顺序。
	return nil // TODO: add indexes to jobs/results and restore order
}

func main() {
	fmt.Println(runOrdered(2, []int{1, 2, 3, 4}))
}
