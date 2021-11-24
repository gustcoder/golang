package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// usando o if init.. variavel fica limitada ao escopo do if/else, não podendo ser usada fora dele
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("outroNumero é maior que zero")
	} else if outroNumero < 0 {
		fmt.Println("outroNumero é menor que zero")
	} else {
		fmt.Println("outroNumero é igual que zero")
	}
}
