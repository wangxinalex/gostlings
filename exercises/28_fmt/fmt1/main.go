// Concept: fmt verbs for structs — %v, %+v, and %#v
// Task: complete formats so it returns the three formatting strings in order
// Expected output: {Ada 36}
// {Name:Ada Age:36}
// main.Person{Name:"Ada", Age:36}
// Hint: fmt.Sprintf("%v", p) prints values; "%+v" adds field names; "%#v" adds the type (Go doc: fmt)

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func formats(p Person) []string {
	// TODO: Return the three formatting strings in order.
	return nil
}

func main() {
	for _, s := range formats(Person{Name: "Ada", Age: 36}) {
		fmt.Println(s)
	}
}
