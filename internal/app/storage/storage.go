package storage

import (
	"fmt"
	"log"
)

type Storage struct {
	storage map[string]*Channel
}

func NewStorage() *Storage {
	return &Storage{
		storage: make(map[string]*Channel),
	}
}

func (s *Storage) Push(key, value string) {
	if _, ok := s.storage[key]; !ok {
		s.storage[key] = NewChannel()
	}
	s.storage[key].Push(value)
	log.Println("в хранилище добавлен ключ", key, "и значение", value)
}

func (s *Storage) Pop(key string, c chan string) {
	if _, ok := s.storage[key]; ok {
		fmt.Println(key)
		c <- s.storage[key].Pop()
	} else {
		c <- ""
	}
}
