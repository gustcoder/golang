package main

import (
	"crud/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// rotas
	router.HandleFunc("/", services.ShowHome).Methods(http.MethodGet)
	router.HandleFunc("/users", services.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/users", services.NewUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", services.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", services.UpdateUser).Methods(http.MethodPut)

	fmt.Println("Listening on port 5000...")
	log.Fatal(http.ListenAndServe(":5000", router))
}
