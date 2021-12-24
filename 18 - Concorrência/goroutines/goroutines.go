package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá Mundo!!!") //goroutine - executa esta funcao mas não precisa esperar ela terminar para ir pra proxima
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
