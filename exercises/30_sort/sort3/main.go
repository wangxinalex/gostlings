// Concept: comparison-based sorting with slices.SortFunc
// Task: complete byAgeDesc so it sorts people by age descending
// Expected output: [{bob 40} {ada 36} {eve 29}]
// Hint: slices.SortFunc(people, func(a, b Person) int { return cmp.Compare(b.Age, a.Age) }) (Go doc: slices).
// Version note: the cmp and slices packages, including slices.SortFunc, were added in Go 1.21.

package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Person struct {
	Name string
	Age  int
}

func byAgeDesc(people []Person) []Person {
	// TODO: Sort people by age descending.
	return people
}

func main() {
	people := []Person{{Name: "ada", Age: 36}, {Name: "bob", Age: 40}, {Name: "eve", Age: 29}}
	fmt.Println(byAgeDesc(people))
}
