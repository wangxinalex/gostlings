// Concept: constraints that require methods, not just underlying types
// Task: define a Labeled constraint and a generic Join function so this program compiles
// Expected output: Ada, Bob
// Hangzhou
// Hint: a constraint can be an ordinary interface with a method; any type that
// implements it works with the generic function (Go Tour: Generics 1)

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Label() string { return p.Name }

type City struct {
	Name string
}

func (c City) Label() string { return c.Name }

// TODO: Define the Labeled interface constraint with a Label() string method.

// TODO: Define Join[T Labeled](items []T, sep string) string that joins every
//       item's label with sep.

func main() {
	fmt.Println(Join([]Person{{Name: "Ada", Age: 36}, {Name: "Bob", Age: 25}}, ", "))
	fmt.Println(Join([]City{{Name: "Hangzhou"}}, ", "))
}
