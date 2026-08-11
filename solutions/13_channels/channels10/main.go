package main

import "fmt"

func trySend(ch chan<- int, value int) bool {
	select {
	case ch <- value:
		return true
	default:
		return false
	}
}

func main() {
	ch := make(chan int, 1)
	fmt.Println(trySend(ch, 7))
}
