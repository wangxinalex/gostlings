// Concept: pass the caller's context through helper layers.
// Task: call source with the exact same context; do not replace it with Background.
// Hint: the helper adds no behavior of its own, so hand the caller's context
//       straight through instead of deriving a new one

package main

import "context"

func lookup(ctx context.Context, source func(context.Context) string) string {
	// TODO: Pass ctx through to source and return its result.
	return ""
}

func main() {}
