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
	// TODO: Extract the "user" value from the context and assign it to user.
	//       If the value is missing, print "no user" and return.
	fmt.Println("no user")
}

func main() {
	ctx := context.Background()
	// TODO: Store the value "Alice" under the key contextKey("user") in the context.
	handler(ctx)
}
