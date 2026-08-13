// Concept: interface embedding composes method sets
// Task: combine two small interfaces into EmployeeInfo and make Employee satisfy it
// Expected behavior: a value assigned to EmployeeInfo exposes both Name and ID
// Hint: embed NameProvider and IDProvider inside EmployeeInfo. Implement Name and ID
//       on Employee. Interface embedding combines method requirements; it does not
//       promote fields or provide implementations.

package main

// NameProvider requires a method that returns a person's name.
type NameProvider interface {
	Name() string
}

// IDProvider requires a method that returns an employee ID.
type IDProvider interface {
	ID() int
}

// EmployeeInfo combines both method sets through interface embedding.
type EmployeeInfo interface {
	NameProvider
	IDProvider
}

type Employee struct {
	name string
	id   int
}

// TODO: Define Name() string and ID() int on Employee.
// The compile-time assertion below should pass only after both methods exist.
var _ EmployeeInfo = Employee{}

func main() {}
