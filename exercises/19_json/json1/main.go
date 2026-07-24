// Concept: JSON marshaling with struct tags
// Task: add json struct tags so Marshal outputs lowercase field names as shown below
// Expected output: {"name":"Alice","age":30}
// Hint: `json:"name"` in a struct tag tells encoding/json what key name to use (Go doc: encoding/json)

package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string // TODO: Add a json tag to rename this field to "name".
	Age  int    // TODO: Add a json tag to rename this field to "age".
}

func main() {
	p := Person{Name: "Alice", Age: 30}
	b, _ := json.Marshal(p)
	fmt.Println(string(b))
}
