package main

import "fmt"

func tryReceive(ch <-chan int) string {
	select {
	case value := <-ch:
		return fmt.Sprintf("received: %d", value)
	default:
		return "no value"
	}
}

func main() {
	fmt.Println(tryReceive(make(chan int)))
}
