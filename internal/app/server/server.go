package server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/kutyrov/my_server_go/internal/app/storage"
)

type Server struct {
	port    string
	storage *storage.Storage
	router  *mux.Router
}

func New(port string) *Server {
	return &Server{
		port:    port,
		storage: storage.NewStorage(),
		router:  mux.NewRouter(),
	}
}

func (s *Server) Start() error {
	//fmt.Println("пишу из старта", s.port)
	my_mux := http.NewServeMux()
	my_mux.HandleFunc("/", s.SwitchMethod)
	//return http.ListenAndServe(s.port, NewHandle("/hi"))
	return http.ListenAndServe(s.port, my_mux)
}

func NewHandle(handle string) *http.ServeMux {
	my_mux := http.NewServeMux()
	my_mux.HandleFunc(handle, home)
	return my_mux
}

func home(w http.ResponseWriter, r *http.Request) {

	//w.Write([]byte("Моя домашняя страница!"))
	//fmt.Println("i am alive", r.Method)
	//fmt.Println(r)
}

func (s *Server) SwitchMethod(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("есть get запрос")
		c := make(chan string)
		go s.PopElem(r.URL.String(), c)
		value := <-c // не считается пока в канал не запишем что нибудь
		// value := s.PopElem(r.URL.String())
		if value == "" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.Write([]byte(value))
		}
	} else if r.Method == "PUT" {
		//fmt.Println("есть post запрос")
		s.PushElem(r.URL.String())
	}
	//fmt.Println(s.storage)
}

func (s *Server) PopElem(data string, c chan string) {
	// добавить обработку таймера
	words := strings.Split(data, "?timeout=")
	key := words[0]
	value := s.storage.Pop(key)
	if value == "" && len(words) == 2 {
		timeout, _ := strconv.Atoi(words[1])
		time.Sleep(time.Duration(timeout) * time.Second)
		value = s.storage.Pop(key)
	}
	c <- value
	// c <- s.storage.Pop(key)
	// value := s.storage.Pop(key)
	// if len(words) == 2 {
	// 	timeout, _ := strconv.Atoi(words[1])
	// 	//duration := timeout * time.Second

	// 	if value == "" {
	// 		timer := time.AfterFunc(time.Duration(timeout)*time.Second, func() {
	// 			value = s.storage.Pop(value)
	// 		})
	// 		defer timer.Stop()
	// 	}
	// }
	// return value

}

func (s *Server) PushElem(data string) {
	words := strings.Split(data, "?v=")
	key := words[0]
	value := words[1]
	s.storage.Push(key, value)
}
