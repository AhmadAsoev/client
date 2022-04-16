package main

import (
	"client/cmd/handleFunc"
	"client/pkg/application/repository"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	GET  = "GET"
	POST = "POST"
)

func main() {

	fmt.Println("server connected")

	if err := repository.Connect(); err != nil {
		log.Fatal("Server can't connect to db: ", err)
	}

	router := mux.NewRouter()
	port := ":7777"

	//Health check the server
	router.HandleFunc("/health", handleFunc.Health).Methods(GET)

	//AddClient adding client into db
	router.HandleFunc("/addClient", handleFunc.AddClient).Methods(POST)

	log.Println("server is running")
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal("Server is not ready!")
	}
}
