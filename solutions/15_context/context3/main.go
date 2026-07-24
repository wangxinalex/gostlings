// Concept: context.WithValue for request-scoped values
// Task: insert a "user" key into the context and extract it in the handler
// Expected output: user: Alice
// Hint: context.WithValue(parent, key, val) stores; ctx.Value(key) retrieves (Go doc: context)

package main

import (
	"context"
	"fmt"
)

type contextKey string

func handler(ctx context.Context) {
	user, ok := ctx.Value(contextKey("user")).(string)
	if !ok {
		fmt.Println("no user")
		return
	}
	fmt.Println("user:", user)
}

func main() {
	ctx := context.WithValue(context.Background(), contextKey("user"), "Alice")
	handler(ctx)
}
