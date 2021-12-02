package main

import "fmt"

func closure() func() {
	texto := "Dentro da função closure"

	return func() {
		fmt.Println(texto)
	}
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	closure := closure()
	closure()
}
