// Concept: struct embedding for composition — Go's alternative to inheritance
// Task: define a Person struct (Name string); embed it in Employee so the program compiles
// Expected output: Employee: Alice (ID: 101)
// Hint: embedding a type without a field name promotes its fields and methods to the outer struct (Go Tour: Moretypes 2-3)

package main

import "fmt"

// TODO: Define a Person struct with a Name string field.

type Employee struct {
	// TODO: Embed Person here (no field name, just the type).
	ID int
}

func main() {
	e := Employee{Person: Person{Name: "Alice"}, ID: 101}
	fmt.Printf("Employee: %s (ID: %d)\n", e.Name, e.ID)
}
