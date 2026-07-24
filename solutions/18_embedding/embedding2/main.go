// Concept: method promotion through embedding
// Task: the GetName method on Person is reached through Employee because Person is embedded; the program calls it but it's missing — define it
// Expected output: Alice (ID: 101)
// Hint: a method defined on an embedded type is promoted to the outer type automatically (Effective Go: Embedding)

package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) GetName() string {
	return p.Name
}

type Employee struct {
	Person
	ID int
}

func (e Employee) String() string {
	return fmt.Sprintf("%s (ID: %d)", e.GetName(), e.ID)
}

func main() {
	e := Employee{Person: Person{Name: "Alice"}, ID: 101}
	fmt.Println(e.String())
}
