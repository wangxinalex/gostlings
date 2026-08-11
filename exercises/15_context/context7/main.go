// Concept: canceling a child context does not cancel its parent or siblings.
// Task: derive and cancel one child, then return both contexts.
// Hint: use context.WithCancel(parent) and defer the returned cancel function.
package main

import "context"

func cancelChild(parent context.Context) (context.Context, context.Context) {
	// TODO: Derive a child from parent and cancel only that child.
	return parent, parent
}

func main() {}
