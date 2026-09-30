// Concept: generic types whose type parameter must be comparable
// Task: define a Set type with Add, Contains, and Len so this program compiles and runs
// Expected output: true
// false
// 2
// Hint: map keys must be comparable, so constrain the type parameter with
// comparable. Back the Set with a map[T]struct{} and give Add a pointer receiver,
// because the map it creates lazily has to survive the call for the other
// methods to see it (Go Tour: Generics 2)

package main

import "fmt"

type Set[T comparable] struct {
	items map[T]struct{}
}

func (s *Set[T]) Add(v T) {
	if s.items == nil {
		s.items = make(map[T]struct{})
	}
	s.items[v] = struct{}{}
}

func (s *Set[T]) Contains(v T) bool {
	_, ok := s.items[v]
	return ok
}

func (s *Set[T]) Len() int {
	return len(s.items)
}

func main() {
	set := Set[string]{}
	set.Add("go")
	set.Add("go")
	set.Add("rust")
	fmt.Println(set.Contains("go"))
	fmt.Println(set.Contains("python"))
	fmt.Println(set.Len())
}
