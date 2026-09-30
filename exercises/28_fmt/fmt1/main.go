// Concept: fmt verbs for structs — %v, %+v, and %#v
// Task: complete formats so it returns the three formatting strings in order
// Expected output: {Ada 36}
// {Name:Ada Age:36}
// main.Person{Name:"Ada", Age:36}
// Hint: the same value has three levels of detail: plain values, values with
//       field names, and a Go-syntax literal that includes the type (Go doc: fmt)
// Stuck?: the %v verb with the + and # modifiers, in that order.

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
