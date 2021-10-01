package server

import (
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/kutyrov/my_server_go/internal/app/storage"
)

type MyStorage interface {
	Push(string, string)
	Pop(string) chan string
}

type Server struct {
	port    string
	storage MyStorage
	router  *mux.Router
	//storage *storage.Storage
}

func New(port string) *Server {
	return &Server{
		port:    port,
		storage: storage.NewStorage(),
		router:  mux.NewRouter(),
	}
}

func (s *Server) Start() error {
	//my_mux := http.NewServeMux()
	//my_mux.HandleFunc("/", s.SwitchMethod)
	s.configureRouter()
	return http.ListenAndServe(s.port, s.router)
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/", s.HandleRoot())
}

func (s *Server) HandleRoot() http.HandlerFunc {
	return http.HandlerFunc(s.SwitchMethod)
	// func(w http.ResponseWriter, r *http.Request) {
	// 	// io.WriteString(w, "hello")
	// }
}

func (s *Server) SwitchMethod(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 	c := make(chan string)
		// 	go s.PopElem(r.URL.String(), c)
		// 	value := <-c // значение не считается пока в канал не запишем что нибудь
		// 	if value == "" {
		// 		w.WriteHeader(http.StatusNotFound)
		// 	} else {
		// 		w.Write([]byte(value))
		// 	}
		// } else if r.Method == "PUT" {
		var key string
		var timeout int
		// парсим таймаут
		s.PopElem(key, time.Duration(timeout))
	} else if r.Method == "PUT" {
		s.PushElem(r.URL.String())
	}
}

func (s *Server) PopElem(key string, timeout time.Duration) string {
	// words := strings.Split(data, "?timeout=")
	// key := words[0]
	// value := s.storage.Pop(key)
	// if value == "" && len(words) == 2 {
	// 	timeout, _ := strconv.Atoi(words[1])
	// 	time.Sleep(time.Duration(timeout) * time.Second)
	// 	value = s.storage.Pop(key)
	// }
	// c <- value
	select {
	case res := <-s.storage.Pop(key):
		return res
	case <-time.After(timeout * time.Second):
		return ""
	}
}

func (s *Server) PushElem(data string) {
	words := strings.Split(data, "?v=")
	key := words[0]
	value := words[1]
	s.storage.Push(key, value)
}
