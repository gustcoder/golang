package main

import "fmt"

func main() {
	canal := make(chan string, 2)

	texto := "Olá mundo"
	canal <- texto
	canal <- "Hello world"
	// canal <- "Exceeded buffer" // deadlock

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
