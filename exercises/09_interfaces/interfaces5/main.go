// Concept: type switches dispatch on the dynamic type of an interface value
// Task: use a type switch in describe so each value gets its correct label
// Expected output: 42 is an int
// hello is a string
// 3.14 is a float64
// Hint: switch v := v.(type) { case int: ... }; use fmt.Sprintf to build the
//       strings (Go Tour: Methods 16)

package main

import "fmt"

func describe(v any) string {
	// TODO: Type-switch over v and return:
	//       int:     "<v> is an int"
	//       string:  "<v> is a string"
	//       float64: "<v> is a float64"
	//       default: "unknown"
}

func main() {
	fmt.Println(describe(42))
	fmt.Println(describe("hello"))
	fmt.Println(describe(3.14))
}
