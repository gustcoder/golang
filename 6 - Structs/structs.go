package main

import "fmt"

func main() {
	type usuario struct {
		nome  string
		idade uint16
	}

	var user usuario

	user.nome = "Gustavo"
	user.idade = 32

	fmt.Println(user)

	// outras formas de declarar/inicializar uma struct
	user2 := usuario{}
	user2.nome = "Novo Usuario"
	user2.idade = 50

	fmt.Println(user2)

	user3 := usuario{"User 3", 33}

	fmt.Println(user3)
}
