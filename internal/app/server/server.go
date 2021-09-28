package server

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/kutyrov/my_server_go/internal/app/storage"
)

type Server struct {
	port    string
	storage *storage.Storage
}

func New(port string) *Server {
	return &Server{
		port:    port,
		storage: storage.NewStorage(),
	}
}

func (s *Server) Start() error {
	my_mux := http.NewServeMux()
	my_mux.HandleFunc("/", s.SwitchMethod)
	return http.ListenAndServe(s.port, my_mux)
}

func (s *Server) SwitchMethod(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		c := make(chan string)
		go s.PopElem(r.URL.String(), c)
		value := <-c // значение не считается пока в канал не запишем что нибудь
		if value == "" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.Write([]byte(value))
		}
	} else if r.Method == "PUT" {
		s.PushElem(r.URL.String())
	}
}

func (s *Server) PopElem(data string, c chan string) {
	words := strings.Split(data, "?timeout=")
	key := words[0]
	value := s.storage.Pop(key)
	if value == "" && len(words) == 2 {
		timeout, _ := strconv.Atoi(words[1])
		time.Sleep(time.Duration(timeout) * time.Second)
		value = s.storage.Pop(key)
	}
	c <- value
}

func (s *Server) PushElem(data string) {
	words := strings.Split(data, "?v=")
	key := words[0]
	value := words[1]
	s.storage.Push(key, value)
}
