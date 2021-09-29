package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChannel_Push(t *testing.T) {
	c := NewChannel()
	print(c.queue)
	c.Push("1")
	c.Push("2")
	c.Push("3")
	res1 := <-c.Pop()
	res2 := <-c.Pop()
	res3 := <-c.Pop()
	assert.Equal(t, res1, "1")
	assert.Equal(t, res2, "2")
	assert.Equal(t, res3, "3")

}
