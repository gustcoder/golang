package main

import "fmt"

func main() {
	var array1 [5]string

	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [2]string{"inferencia", "tipo"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice1 := []int{1, 2, 3}

	slice1 = append(slice1, 4)
	fmt.Println(slice1)

	slice2 := array3[0:3] // primeiro param é inclusivo (ponto de partida), segundo parametro é exclusivo, pega ate o anterior a ele... 0, 1 , 2
	fmt.Println(slice2)

	array3[2] = 99 // por referencia, altera a fatia do array
	fmt.Println(slice2)
}
