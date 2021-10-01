package storage

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

}

func (s *Storage) Pop(key string) chan string {
	c := make(chan string)
	c <- s.storage[key].Pop()
	return c
}
