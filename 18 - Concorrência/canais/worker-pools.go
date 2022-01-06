package main

import "fmt"

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	execTarefas(tarefas)
	obtemResultados(resultados)
}

func execTarefas(tarefas chan int) {
	for i := 0; i < 45; i++ {
		tarefas <- i
	}

	close(tarefas)
}

func obtemResultados(resultados chan int) {
	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

// tarefas apenas recebe dados (<-chan), e resultados envia dados (chan<-)
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
