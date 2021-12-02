package main

import "fmt"

func recuperarExecucao() {
	fmt.Println("Tentativa de recuperar a execução")
	// retornara false pois é bool
	if r := recover(); r != nil { // se resultado == nil ele ignora e toca o barco
		fmt.Println("Execução recuperada com sucesso")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()

	media := (n1 + n2) / 2
	fmt.Println(media)

	if media > 6 {
		return true
	}

	if media < 6 {
		return false
	}

	// como uma "exception"
	panic("Média é exatamente 6")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Término")
}
