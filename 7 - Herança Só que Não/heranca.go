package main

import "fmt"

func main() {
	type pessoa struct {
		nome  string
		idade uint16
	}

	type estudante struct {
		pessoa    // aqui só passa a struct, sem tipo, para acessar as props diretamente
		curso     string
		faculdade string
	}

	var p1 pessoa

	p1.nome = "Gustavo"
	p1.idade = 32

	fmt.Println(p1)

	e1 := estudante{p1, "Ciência da Computação", "Unis"}

	fmt.Println(e1.nome)
}
