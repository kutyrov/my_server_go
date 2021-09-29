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
		queue: make(chan string),
	}
}

func (c *Channel) Push(value string) {
	c.mu.Lock()
	c.queue <- value
	c.mu.Unlock()
}

func (c *Channel) Pop() chan string {
	c.mu.Lock()
	temp := make(chan string)
	res := <-c.queue
	c.mu.Unlock()
	temp <- res
	return temp
}
