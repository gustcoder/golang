package main

import "fmt"

func main() {
	canal := escrever("Testando")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(text string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Texto: %s", text)
		}
	}()

	return canal
}
