package main

import (
	"fmt"
)

func main() {
	// i := 0
	// for i < 10 {
	// 	time.Sleep(time.Second)
	// 	fmt.Println("Incrementando i")
	// 	i++
	// }
	// fmt.Println(i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando j", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"Jesus", "Davi", "Paulo"}
	// for indice, nome := range nomes {
	for _, nome := range nomes {
		// fmt.Println(indice, nome)
		fmt.Println(nome)
	}

	for indice, letra := range "DEUS" {
		// fmt.Println(indice, letra)
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{"nome": "Gustavo"}

	for _, valor := range usuario {
		fmt.Println(valor)
	}
}
