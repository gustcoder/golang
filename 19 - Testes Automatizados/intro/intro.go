package main

import (
	"fmt"
	"intro/enderecos"
)

func main() {
	endereco := enderecos.CheckTipoEndereco("rua dos Bobos, numero tal")
	fmt.Println(endereco)
}
