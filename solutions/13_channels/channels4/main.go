package main

import "fmt"

func read(ch <-chan int) (int, bool) {
	value, ok := <-ch
	return value, ok
}

func main() {
	ch := make(chan int)
	close(ch)
	value, ok := read(ch)
	fmt.Println(value, ok)
}
