// Concept: sync.Map is useful for concurrent registries with independent keys.
// Task: implement string-keyed Load, Store, and canonical LoadOrStore operations.
// Hint: keep the zero-value registry usable and return the value that won publication.
package main

import "sync"

type registry struct{ values sync.Map }

func (r *registry) Load(key string) (string, bool) {
	// TODO: Load and type-assert a string value.
	return "", false
}

func (r *registry) Store(key, value string) {
	// TODO: Store a string value.
}

func (r *registry) LoadOrStore(key, value string) (string, bool) {
	// TODO: Publish one canonical value and report whether it already existed.
	return "", false
}

func main() {}
