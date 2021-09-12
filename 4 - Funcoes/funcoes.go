package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

// funcao com mais de 1 retorno
func calculosMatematicos(n1, n2 int) (int, int) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	soma := somar(5, 10)
	fmt.Println(soma)

	var f = func() {
		fmt.Println("Esta é uma função")
	}
	f()

	var p = func(parametro string) string {
		return parametro
	}
	resultado := p("Passando parametro")
	fmt.Println(resultado)

	// funcao com mais de 1 retorno
	resultadoSoma, resultadoSubtracao := calculosMatematicos(20, 30)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	// usando apenas um dos retornos da funcao que possui mais de um retorno
	apenasUmResultado, _ := calculosMatematicos(2, 2)
	fmt.Println(apenasUmResultado)
}
