// Concept: generic types with methods
// Task: define a generic Stack type with Push and pointer-receiver Pop methods
// Expected output: 2
// Hint: back the Stack with a slice; Pop returns the last element (Go Tour: Generics 2)

package main

import "fmt"

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() T {
	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return v
}

func main() {
	s := &Stack[int]{}
	s.Push(1)
	s.Push(2)
	fmt.Println(s.Pop())
}
