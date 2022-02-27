package services

import (
	"crud/db"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func ResponseError(rw http.ResponseWriter, msg string) {
	rw.Write([]byte(msg))
	return
}

// interface{} => como se fosse o mixed, onde pode chegar qualquer tipo
func ResponseOK(rw http.ResponseWriter, msg string, parameters interface{}) {
	if err := json.NewEncoder(rw).Encode(parameters); err != nil {
		ResponseError(rw, msg)
	}
}

func ConnectDatabase() (*sql.DB, string) {
	db, err := db.Connect()
	if err != nil {
		return nil, "Error connecting database"
	}

	return db, ""
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)

	ID, err := strconv.ParseUint(request["id"], 10, 32)
	if err != nil {
		ResponseError(rw, "Error converting ID to int")
	}

	if err := json.NewEncoder(rw).Encode(ID); err != nil {
		ResponseError(rw, "Error getting ID")
	}

	db, err := db.Connect()
	if err != nil {
		ResponseError(rw, "Error connecting database")
	}
	defer db.Close()

	//userInfo, err := db.Query("SELECT * FROM usuarios WHERE id = ?")

	// var user User
	// if err := json.Unmarshal(request, &user); err != nil {
	// 	rw.Write([]byte("Error on Getting Request"))
	// 	return
	// }

	//userInfo, err := db.Query("SELECT * FROM usuarios WHERE id = ?")
}

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()
	if err != nil {
		ResponseError(rw, "Error connecting database")
	}
	defer db.Close()

	usersList, err := db.Query("SELECT * FROM usuarios")
	if err != nil {
		ResponseError(rw, "Error on Select")
	}

	// criando um slice para armazenar objetos dos usuarios
	var users []User

	for usersList.Next() {
		var user User

		if err := usersList.Scan(&user.ID, &user.Nome, &user.Email); err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	// faz o json do slice para retornar usuarios
	// userJson, err := json.Marshal(users)
	// if err != nil {
	// 	log.Fatal("Error creating users list Json")
	// }
	// rw.WriteHeader(http.StatusOK)
	// rw.Write(userJson)

	// jeito correto usando json.NewEncode
	ResponseOK(rw, "Error on encode users to JSON", users)
}

func NewUser(rw http.ResponseWriter, r *http.Request) {
	request, err := ioutil.ReadAll(r.Body)

	if err != nil {
		ResponseError(rw, "Error reading NewUser request body")
	}

	var user User

	if err := json.Unmarshal(request, &user); err != nil {
		ResponseError(rw, "Error converting User to struct")
	}

	db, err := db.Connect()
	if err != nil {
		ResponseError(rw, "Error connecting database")
	}
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO usuarios (nome, email) VALUES (?, ?)")

	if err != nil {
		ResponseError(rw, "Error on Prepare statement")
	}
	defer db.Close()

	insert, err := statement.Exec(user.Nome, user.Email)
	if err != nil {
		ResponseError(rw, "Error on Insert")
	}

	lastUserId, err := insert.LastInsertId()
	if err != nil {
		ResponseError(rw, "Error getting Last Inserted User ID")
	}

	fmt.Println(user) // including response to server
	rw.WriteHeader(http.StatusCreated)
	rw.Write(
		[]byte(
			fmt.Sprintf("User created successfully with ID %d", lastUserId),
		),
	)
}
