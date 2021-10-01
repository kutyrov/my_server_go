package storage

import (
	"sync"
)

type Channel struct {
	mu    sync.Mutex
	queue chan string
	Len   int
}

func NewChannel() *Channel {
	return &Channel{
		queue: make(chan string, 10),
	}
}

func (c *Channel) Push(value string) {
	c.mu.Lock()
	c.queue <- value
	c.Len += 1
	c.mu.Unlock()
}

func (c *Channel) Pop() string {
	c.mu.Lock()
	res := <-c.queue
	c.Len -= 1
	c.mu.Unlock()
	return res
}

func (c *Channel) GetLen() int {
	return c.Len
}

func (c *Channel) SafePop() string {
	var res string
	if c.Len > 0 {
		res = c.Pop()
	}
	return res
}
