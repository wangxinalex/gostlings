// Concept: sync.RWMutex allows concurrent readers and exclusive writers.
// Task: protect a small cache so Get and Put are safe during concurrent use.
// Hint: use RLock/RUnlock for Get and Lock/Unlock for Put; protect the map itself too.
package main

import "sync"

type readWriteCache struct {
	mu     sync.RWMutex
	values map[string]string
}

func (c *readWriteCache) Get(key string) (string, bool) {
	// TODO: Read the cache under a read lock.
	return "", false
}

func (c *readWriteCache) Put(key, value string) {
	// TODO: Initialize and update the cache under the write lock.
}

func main() {}
