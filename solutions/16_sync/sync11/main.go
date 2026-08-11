// Concept: sync.Map is useful for concurrent registries with independent keys.
// Task: implement string-keyed Load, Store, and canonical LoadOrStore operations.
// Hint: keep the zero-value registry usable and return the value that won publication.
package main

import "sync"

type registry struct{ values sync.Map }

func (r *registry) Load(key string) (string, bool) {
	value, ok := r.values.Load(key)
	if !ok {
		return "", false
	}
	text, ok := value.(string)
	return text, ok
}

func (r *registry) Store(key, value string) { r.values.Store(key, value) }

func (r *registry) LoadOrStore(key, value string) (string, bool) {
	actual, loaded := r.values.LoadOrStore(key, value)
	return actual.(string), loaded
}

func main() {}
