package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/michalnov/Stockpile-Server/build/handler"
)

func main() {
	fmt.Println("Hello server")
	router := mux.NewRouter()
	fmt.Println("Run server on port: 3311")
	http.Handle("/", router)
	router.HandleFunc("/what", handler.Hello_Handler).Methods("GET")
	router.HandleFunc("/login", handler.LoginHandler).Methods("POST")
	router.HandleFunc("/update", handler.UpdateStock).Methods("POST")
	http.ListenAndServe(":3311", router)
}
