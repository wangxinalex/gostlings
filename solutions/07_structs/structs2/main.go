// Concept: initializing structs with field names
// Task: fix this program so it compiles and runs
// Expected output: {18 Tom}
// Hint: a struct literal can name its fields in any order: Person{Name: "Tom", Age: 18} (Go Tour: Moretypes 5)

package main

import "fmt"

type Person struct {
	Age  int
	Name string
}

func main() {
	p := Person{Name: "Tom", Age: 18}
	fmt.Println(p)
}
