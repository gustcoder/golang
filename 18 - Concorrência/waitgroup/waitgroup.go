package main

import (
	"fmt"
	"time"
)

func main() {
	escrever("Olá Mundo!!!") //goroutine - executa esta funcao mas não precisa esperar ela terminar para ir pra proxima
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
