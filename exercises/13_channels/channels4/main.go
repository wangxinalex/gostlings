// Concept: comma-ok receive distinguishes a closed channel from a real zero value
// Task: use the receive form that reports whether a value was actually received
// Expected behavior: a closed channel reports value 0 with ok=false
// Hint: use value, ok := <-ch; ok is false only after the channel is closed and drained

package main

import "fmt"

func read(ch <-chan int) (int, bool) {
	// 思路：接收 int 时，零值 0 可能是真实数据，也可能是关闭后的默认值；
	// comma-ok 的第二个返回值才是判断 channel 生命周期的依据。
	value := <-ch
	return value, true // TODO: receive with comma-ok and return the real status
}

func main() {
	ch := make(chan int)
	close(ch)
	value, ok := read(ch)
	fmt.Println(value, ok)
}
