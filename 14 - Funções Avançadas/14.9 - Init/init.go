package main

import "fmt"

var n int

func init() {
	fmt.Println("Executando antes da função main")
	n = 10
}

func main() {
	fmt.Println("Executando função main")
	fmt.Println(n)
}
