package storage

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStorage_Push(t *testing.T) {
	s := NewStorage()
	print(s.storage)
	s.Push("first", "1")
	res := s.Pop("first", time.Duration(0))
	assert.Equal(t, res, "1")
}
