package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplex(escrever("Canal 1"), escrever("Canal 2"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplex(canal1 <-chan string, canal2 <-chan string) <-chan string {
	canalSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canal1:
				canalSaida <- mensagem
			case mensagem := <-canal2:
				canalSaida <- mensagem
			}
		}
	}()

	return canalSaida
}

func escrever(text string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Texto: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
