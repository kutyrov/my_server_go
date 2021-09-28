package storage

import (
	"sync"
)

type Storage struct {
	mu      sync.Mutex
	storage map[string][]string
}

func NewStorage() *Storage {
	return &Storage{
		mu:      sync.Mutex{},
		storage: make(map[string][]string),
	}
}

func (s *Storage) Push(key, value string) {
	s.mu.Lock()

	s.storage[key] = append(s.storage[key], value)

	s.mu.Unlock()

}

func (s *Storage) Pop(key string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.storage[key]) != 0 {
		//s.mu.Lock()
		temp := s.storage[key][0]
		s.storage[key] = s.storage[key][1:]
		//s.mu.Unlock()
		return temp
	}

	return ""
}
