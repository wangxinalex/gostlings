// Concept: defining methods with a value receiver
// Task: define the Area method so this program compiles and runs
// Expected output: 6
// Hint: a method is a function with a receiver argument before the method name (Go Tour: Methods 1)

package main

import "fmt"

type Rectangle struct {
	W, H int
}

func (r Rectangle) Area() int {
	return r.W * r.H
}

func main() {
	r := Rectangle{2, 3}
	fmt.Println(r.Area())
}
