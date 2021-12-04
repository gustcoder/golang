package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	cargo string
}

func (u usuario) salvar() {
	//fmt.Println("Usuario " + u.nome + " salvo com sucesso")
	fmt.Printf("Usuario %s salvo com sucesso!\n", u.nome)
}

func (u usuario) possuiCargo() bool {
	return u.cargo != ""
}

func (u *usuario) setCargo(cargo string) {
	u.cargo = cargo
}

func main() {
	usuario1 := usuario{nome: "Gustavo Ramos", idade: 32, cargo: "Engenheiro de Software"}
	usuario1.salvar()
	fmt.Println(usuario1.possuiCargo())

	usuario2 := usuario{"Diego", 37, ""}
	usuario2.salvar()
	fmt.Println(usuario2.possuiCargo())

	usuario2.setCargo("Youtuber")
	fmt.Println(usuario2)
}
