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
	router.HandleFunc("/users", services.NewUser).Methods(http.MethodPost)

	fmt.Println("Listening on port 5000...")
	log.Fatal(http.ListenAndServe(":5000", router))
}
