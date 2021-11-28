package main

import "fmt"

func main() {
	retorno := func(parametro string) string {
		return fmt.Sprintf("Recebido -> %s", parametro)
	}("Passando Parametro")
	fmt.Println(retorno)
}
