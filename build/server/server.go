package server

import (
	//"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/michalnov/Stockpile-Server/build/handler"
)

type Server struct {
	conf struct {
		port string
	}
	Terminate chan int
	router    *mux.Router
}

func (s *Server) routes() {
	s.router.HandleFunc("/hello", handler.Hello_Handler).Methods("GET")
	//s.router.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
	//	s.Terminate <- 0
	//})
}

func (s *Server) Start() {
	fmt.Println("Run server on port: " + s.conf.port)
	http.Handle("/", s.router)
	s.router.HandleFunc("/hello", handler.Hello_Handler).Methods("GET")
	s.router.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		s.Terminate <- 0
	})
	http.ListenAndServe(s.conf.port, s.router)
}

func NewServer() (Server, error) {
	out := Server{}
	out.router = mux.NewRouter()
	return out, nil
}

func (s *Server) SetPort(port string) {
	s.conf.port = port
}
