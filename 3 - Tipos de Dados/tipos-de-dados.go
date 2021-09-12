package main

import (
	"errors"
	"fmt"
)

func main() {
	// INTEIROS
	var numero int8 = 100
	fmt.Println(numero)

	// somente int assume a arquitetura do computador (ex.: 32bits, 64 bits)
	var numero1 int = 100000
	fmt.Println(numero1)

	var apenasPositivo uint = 100
	fmt.Println(apenasPositivo)

	// REAIS
	var real float32 = 999.10
	fmt.Println(real)

	// para float melhor declarar assim
	floatPorInferencia := 455562.25
	fmt.Println(floatPorInferencia)

	// STRING

	var nome string = "Gustavo"
	sobrenome := "Ramos"

	fmt.Println(nome, sobrenome)

	char := "A"
	fmt.Println(char)

	// BOOLEAN

	var booleano bool = true
	fmt.Println(booleano)

	// TIPO ERROR DO GO
	var erro error
	fmt.Println(erro)

	var customError error = errors.New("Este é um erro customizado")
	fmt.Println(customError)
}
