// Concept: the fmt.Stringer interface controls how values are printed
// Task: this program should print 25°C, but it prints 25; define a String method on Celsius
// Expected output: 25°C
// Hint: fmt looks for a String() string method on the value it prints (Go Tour: Methods 17)

package main

import "fmt"

type Celsius float64

// TODO: Define a String method on Celsius returning fmt.Sprintf("%g°C", float64(c)).

func main() {
	fmt.Println(Celsius(25))
}
