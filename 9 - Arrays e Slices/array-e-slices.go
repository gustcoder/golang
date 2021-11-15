package main

import "fmt"

func main() {
	var array1 [5]string

	// declaracao padrao
	array1[0] = "Posição 1"
	fmt.Println(array1)

	// declaracao por inferencia de  tipo
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

	// Arrays Internos (tipo, tamanho, capacidade maxima, que é opcional)
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	// se tentar extrapolar a capcidade, o Go
	// se encarrega de "dobrar" a capacidade do array
	// interno, fazendo o slice ter capacidade ilimitada
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]int)
	fmt.Println(slice4)

	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
