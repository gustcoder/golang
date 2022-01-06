package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá mundo", canal)

	// for {
	// 	mensagem, canalAberto := <-canal
	// 	if !canalAberto {
	// 		break
	// 	}

	// 	fmt.Println(mensagem)
	// }

	for mensagem := range canal {
		fmt.Println(mensagem)
	}
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}
