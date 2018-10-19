package main

import (
	"fmt"

	"github.com/michalnov/Stockpile-Server/build/server"
)

func main() {
	fmt.Println("Hello server")

	term := make(chan int)

	serv, _ := server.NewServer()
	//go runServ(serv, term, ":8080")
	go runServ(serv, term, ":8080")
	_ = <-term
}

func runServ(s server.Server, terminate chan int, port string) {
	s.Terminate = terminate
	s.SetPort(port)
	s.Start()
}
