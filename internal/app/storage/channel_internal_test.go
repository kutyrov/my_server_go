package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChannel_Push(t *testing.T) {
	c := NewChannel()
	c.Push("1")
	c.Push("2")
	c.Push("3")
	res1 := c.Pop()
	res2 := c.Pop()
	res3 := c.Pop()
	//res4 := c.SafePop()
	assert.Equal(t, res1, "1")
	assert.Equal(t, res2, "2")
	assert.Equal(t, res3, "3")
	//assert.Equal(t, res4, "")

}
