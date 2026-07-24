// Concept: recover catches a panic inside a deferred function
// Task: add a deferred recover so the program prints the recovered message instead of crashing
// Expected output: recovered: something went wrong
// Hint: recover() returns the panic value; you must call it inside a deferred function (Go doc: builtin)

package main

import "fmt"

func mayPanic() {
	panic("something went wrong")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered:", r)
		}
	}()
	mayPanic()
	fmt.Println("this line never runs if there is no recover")
}
