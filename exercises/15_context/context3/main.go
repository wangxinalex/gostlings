// Concept: context.WithValue carries request-scoped values with typed keys.
// Task: look up the user value and return the documented fallback when absent.
// Hint: use a private key type, then use the comma-ok type assertion on ctx.Value.
package main

import (
	"context"
	"fmt"
)

type userKey struct{}

func handler(ctx context.Context) string {
	// TODO: Read userKey{} as a string and return "user: guest" when it is missing.
	return "user: guest"
}

func main() {
	ctx := context.WithValue(context.Background(), userKey{}, "Alice")
	fmt.Println(handler(ctx))
}
