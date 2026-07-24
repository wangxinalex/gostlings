// Concept: accessing fields of nested structs
// Task: initialize a Person whose city is "Hangzhou" and print that city
// Expected output: Hangzhou
// Hint: struct fields are accessed with a dot, and nested fields chain dots: p.Addr.City (Go Tour: Moretypes 2-3)

package main

import "fmt"

type Address struct {
	City string
}

type Person struct {
	Name string
	Addr Address
}

func main() {
	p := Person{Name: "Bob", Addr: Address{City: "Hangzhou"}}
	fmt.Println(p.Addr.City)
}
