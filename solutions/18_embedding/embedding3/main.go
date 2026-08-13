// Concept: interface embedding composes method sets
// Task: combine two small interfaces into EmployeeInfo and make Employee satisfy it
// Expected behavior: a value assigned to EmployeeInfo exposes both Name and ID
// Hint: embed NameProvider and IDProvider inside EmployeeInfo. Implement Name and ID
//       on Employee. Interface embedding combines method requirements; it does not
//       promote fields or provide implementations.

package main

type NameProvider interface {
	Name() string
}

type IDProvider interface {
	ID() int
}

type EmployeeInfo interface {
	NameProvider
	IDProvider
}

type Employee struct {
	name string
	id   int
}

func (e Employee) Name() string {
	return e.name
}

func (e Employee) ID() int {
	return e.id
}

var _ EmployeeInfo = Employee{}

func main() {}
