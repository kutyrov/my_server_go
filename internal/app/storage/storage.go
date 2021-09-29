package storage

import "time"

type Storage struct {
	storage map[string]*Channel
}

func NewStorage() *Storage {
	return &Storage{
		storage: make(map[string]*Channel),
	}
}

func (s *Storage) Push(key, value string) {
	if _, err := s.storage[key]; err {
		s.storage[key] = NewChannel()
	}
	s.storage[key].Push(value)

}

func (s *Storage) Pop(key string, timeout time.Duration) string {
	select {
	case res := <-s.storage[key].Pop():
		return res
	case <-time.After(timeout):
		return ""
	}
}
