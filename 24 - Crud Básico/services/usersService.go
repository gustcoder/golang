package services

import (
	"crud/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func NewUser(rw http.ResponseWriter, r *http.Request) {
	request, err := ioutil.ReadAll(r.Body)

	if err != nil {
		rw.Write([]byte("Error reading NewUser request body"))

		return
	}

	var user User

	if err := json.Unmarshal(request, &user); err != nil {
		rw.Write([]byte("Error converting User to struct"))

		return
	}

	db, err := db.Connect()
	if err != nil {
		rw.Write([]byte("Error connecting database"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO usuarios (nome, email) VALUES (?, ?)")

	if err != nil {
		rw.Write([]byte("Error on Prepare statement"))
		return
	}
	defer db.Close()

	insert, err := statement.Exec(user.Nome, user.Email)
	if err != nil {
		rw.Write([]byte("Error on Insert"))
		return
	}

	lastUserId, err := insert.LastInsertId()
	if err != nil {
		rw.Write([]byte("Error getting Last Inserted User ID"))
		return
	}

	fmt.Println(user) // including response to server
	rw.WriteHeader(http.StatusCreated)
	rw.Write(
		[]byte(
			fmt.Sprintf("User created successfully with ID %d", lastUserId),
		),
	)
}
