// Concept: interfaces and implementing them implicitly
// Task: define the Speak method on Dog so this program compiles and runs
// Expected output: Woof!
// Hint: a type implements an interface simply by defining its methods (Go Tour: Methods 9-10)

package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

func announce(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	announce(Dog{})
}
