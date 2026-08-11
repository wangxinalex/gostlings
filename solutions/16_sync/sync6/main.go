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
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.values[key]
	return value, ok
}

func (c *readWriteCache) Put(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.values == nil {
		c.values = make(map[string]string)
	}
	c.values[key] = value
}

func main() {}
