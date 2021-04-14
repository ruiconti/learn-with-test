package main

import "sync"

type Counter struct {
	mutex    sync.Mutex
	register int
}

func (c *Counter) Value() int {
	return c.register
}

func (c *Counter) Inc() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.register++
}
