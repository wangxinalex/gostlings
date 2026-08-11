// Concept: pass the caller's context through helper layers.
// Task: call source with the exact same context; do not replace it with Background.
// Hint: this helper is intentionally small: return source(ctx).
package main

import "context"

func lookup(ctx context.Context, source func(context.Context) string) string {
	return source(ctx)
}

func main() {}
