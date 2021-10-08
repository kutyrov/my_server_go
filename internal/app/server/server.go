package server

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/kutyrov/my_server_go/internal/app/storage"
)

type MyStorage interface {
	Push(string, string)
	Pop(string, chan string)
}

type Server struct {
	port    string
	storage MyStorage
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
	//s.configureRouter()
	//return http.ListenAndServe(s.port, s.router)
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(s.HandleRoot()))
	return http.ListenAndServe(s.port, mux)
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/", s.HandleRoot())
}

func (s *Server) HandleRoot() http.HandlerFunc {
	return http.HandlerFunc(s.SwitchMethod)
}

func (s *Server) SwitchMethod(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// если url не соответствует шаблону то вернем фигню
		key := r.URL.Path[1:] // убрали первый слэш

		var timeout int
		if _, ok := r.URL.Query()["timeout"]; ok {
			var err error
			timeout, err = strconv.Atoi(r.URL.Query()["timeout"][0])
			if err != nil {
				log.Println("Не получилось спарсить timeout")
				// вернём фигню
			}
		}

		log.Println("Таймаут", timeout)

		// парсим таймаут
		res := s.PopElem(key, time.Duration(timeout))
		log.Println("Отдаём значение", res)
		// здесь формируем response
		w.Write([]byte(res))
	} else if r.Method == "PUT" {
		// если url не соответствует шаблону то вернем фигню
		key := r.URL.Path[1:] // убрали первый слэш
		value := r.URL.Query()["v"][0]
		s.PushElem(key, value)
	}
}

func (s *Server) PopElem(key string, timeout time.Duration) string {
	c := make(chan string)
	//quit := make(chan string)
	go s.storage.Pop(key, c)
	select {
	case res := <-c:
		return res
	case <-time.After(timeout * time.Second):
		return ""
	}
}

func (s *Server) PushElem(key, value string) {
	s.storage.Push(key, value)
}
