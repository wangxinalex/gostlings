// Concept: closures capture variables from the enclosing scope
// Task: implement newCounter so each returned function keeps its own counter,
//       increments it, and returns the new value
// Expected output: 3
// 2
// Hint: return func() int { ... }; the counter variable lives in the closure's
//       captured scope, and each newCounter call gets its own copy
//       (builds on Go Tour: Basics 4-7; closures are not covered in the Tour)

package main

import "fmt"

func newCounter() func() int {
	// TODO: Return a closure that increments a captured counter and returns it.
	//       Each call to newCounter must return an independent counter.
}

func main() {
	a := newCounter()
	b := newCounter()
	a()
	a()
	b()
	fmt.Println(a())
	fmt.Println(b())
}
