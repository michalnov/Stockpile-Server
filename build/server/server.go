package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	conf struct {
		port string
	}
	router *mux.Router
}

func (s *Server) routes(){
	s.router.
}

func (s *Server) Start() {
	http.ListenAndServe(s.conf.port, s.router)
}

func NewServer() (Server, error) {
	out := Server
	out.router = mux.NewRouter()
}

func (s *Server) SetPort(port int) {
	s.conf.port = ":" + string(port)
}
