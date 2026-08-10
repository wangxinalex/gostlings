// Concept: fan-out — several workers consume one jobs channel
// Task: square every job and close results after every worker exits
// Expected behavior: one squared result per job, followed by a closed output
// Hint: workers range over jobs; a coordinator waits for them before closing results

package main

import "fmt"

func squareWorkers(workers int, jobs <-chan int) <-chan int {
	// 思路：jobs 的关闭告诉所有 worker 没有新任务；results 的关闭必须晚于所有 worker，
	// 否则仍在发送结果的 worker 会向已关闭 channel 发送而 panic。
	return nil // TODO: start workers and coordinate result closure
}

func main() {
	jobs := make(chan int, 3)
	jobs <- 1
	jobs <- 2
	jobs <- 3
	close(jobs)
	for result := range squareWorkers(2, jobs) {
		fmt.Println(result)
	}
}
