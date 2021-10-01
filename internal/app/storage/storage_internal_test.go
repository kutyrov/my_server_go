package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage_PushandPop(t *testing.T) {
	s := NewStorage()
	s.Push("first", "1")
	s.Push("first", "2")
	s.Push("first", "3")
	res1, res2, res3 := <-s.Pop("first"), <-s.Pop("first"), <-s.Pop("first")

	assert.Equal(t, res1, "1")
	assert.Equal(t, res2, "2")
	assert.Equal(t, res3, "3")

}

func TestStorage_SomeChannels(t *testing.T) {
	s := NewStorage()
	s.Push("first", "1")
	s.Push("first", "2")
	s.Push("second", "one")
	s.Push("second", "two")
	res1, res2 := <-s.Pop("first"), <-s.Pop("first")
	res3, res4 := <-s.Pop("second"), <-s.Pop("second")
	assert.Equal(t, res1, "1")
	assert.Equal(t, res2, "2")
	assert.Equal(t, res3, "one")
	assert.Equal(t, res4, "two")
}
