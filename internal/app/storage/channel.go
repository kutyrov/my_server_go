package storage

import (
	"sync"
)

type Channel struct {
	mu    sync.Mutex
	queue chan string
}

func NewChannel() *Channel {
	return &Channel{
		queue: make(chan string, 10),
	}
}

func (c *Channel) Push(value string) {
	c.mu.Lock()
	c.queue <- value
	c.mu.Unlock()
}

func (c *Channel) Pop() string {
	c.mu.Lock()
	res := <-c.queue
	c.mu.Unlock()
	return res
}
