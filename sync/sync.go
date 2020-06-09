package sync

import "sync"

type Counter struct {
	// here notice we do not simply embed sync.Mutex
	// because we do not want the methods of sync.Mutex becomes part of our struct
	mu sync.Mutex
	value int
}

// return the pointer to the new counter
func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
