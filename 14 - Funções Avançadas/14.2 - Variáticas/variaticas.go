package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func boasVindas(texto string, nomes ...string) {
	for _, nome := range nomes {
		fmt.Println(texto, nome)
	}
}

func main() {
	soma := soma(10, 2, 3, 5)
	fmt.Println(soma)

	boasVindas("Bem-vindo(a)", "Gustavo", "Phil", "Sueli", "Guilherme")
}
