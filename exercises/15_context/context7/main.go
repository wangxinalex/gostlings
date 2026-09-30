// Concept: canceling a child context does not cancel its parent or siblings.
// Task: derive and cancel one child, then return both contexts.
// Hint: derive the child from the parent, keep its cancel function, and release
//       only the child; the parent and any siblings must keep running.
// Stuck?: context.WithCancel with the parent as its argument.

package main

import "context"

func cancelChild(parent context.Context) (context.Context, context.Context) {
	// TODO: Derive a child from parent and cancel only that child.
	return parent, parent
}

func main() {}
