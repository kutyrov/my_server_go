package storage

import (
	"fmt"
	"sync"
)

type Channel struct {
	mu    sync.Mutex
	queue chan string
}

func NewChannel() *Channel {
	return &Channel{
		queue: make(chan string, 100),
	}
}

func (c *Channel) Push(value string) {
	c.mu.Lock()
	fmt.Println("push locked")
	c.queue <- value
	c.mu.Unlock()
	fmt.Println("push unlocked")
}

func (c *Channel) Pop() string {
	for {
		if len(c.queue) > 0 {
			c.mu.Lock()
			fmt.Println("pop locked")
			res := <-c.queue
			c.mu.Unlock()
			fmt.Println("pop unlocked")
			return res
		}
	}

}
