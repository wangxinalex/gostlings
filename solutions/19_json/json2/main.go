// Concept: JSON unmarshaling into a struct
// Task: declare a Person variable and Unmarshal the JSON bytes into it, then print the name
// Expected output: Name: Bob, Age: 25
// Hint: json.Unmarshal takes a []byte and a pointer to the target value (Go doc: encoding/json)

package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	data := []byte(`{"name":"Bob","age":25}`)
	var p Person
	json.Unmarshal(data, &p)
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}
