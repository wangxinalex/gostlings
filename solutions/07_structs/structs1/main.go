// Concept: defining a struct type
// Task: define the Person struct so this program compiles and runs
// Expected output: {Alice 18}
// Hint: a struct groups fields together: type Person struct { Name string; Age int } (Go Tour: Moretypes 2)

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Alice", Age: 18}
	fmt.Println(p)
}
