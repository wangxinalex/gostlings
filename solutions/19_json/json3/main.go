// Concept: omitempty tag and handling missing JSON fields
// Task: the Age field should be omitted when it is 0; add the omitempty option to the json tag
// Expected output: {"name":"Carol"}
// Hint: `json:"age,omitempty"` skips the field in the output when it has its zero value (Go doc: encoding/json)

package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}

func main() {
	p := Person{Name: "Carol"}
	b, _ := json.Marshal(p)
	fmt.Println(string(b))
}
