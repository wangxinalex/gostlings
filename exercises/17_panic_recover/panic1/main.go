// Concept: recover catches a panic inside a deferred function
// Task: add a deferred recover so the program prints the recovered message instead of crashing
// Expected output: recovered: something went wrong
// Hint: functions6 introduced function literals. recover() returns the panic
//       value; call it inside a deferred function (Go doc: builtin)

package main

import "fmt"

func mayPanic() {
	panic("something went wrong")
}

func main() {
	// TODO: Defer a function that calls recover() and prints the recovered message.
	mayPanic()
	fmt.Println("this line never runs if there is no recover")
}
