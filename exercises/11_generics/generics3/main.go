// Concept: generic types with methods
// Task: define a generic Stack type with Push and pointer-receiver Pop methods
// Expected output: 2
// Hint: back the Stack with a slice; Pop returns the last element (Go Tour: Generics 2)

package main

import "fmt"

// TODO: Define Stack[T any] and its Push(v T) and Pop() T methods.
//       Push appends to the slice; Pop removes and returns the last element.

func main() {
	s := &Stack[int]{}
	s.Push(1)
	s.Push(2)
	fmt.Println(s.Pop())
}
