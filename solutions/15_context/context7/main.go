// Concept: canceling a child context does not cancel its parent or siblings.
// Task: derive and cancel one child, then return both contexts.
// Hint: use context.WithCancel(parent) and defer the returned cancel function.
package main

import "context"

func cancelChild(parent context.Context) (context.Context, context.Context) {
	child, cancel := context.WithCancel(parent)
	defer cancel()
	return parent, child
}

func main() {}
