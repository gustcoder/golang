package auxiliar

import (
	"fmt"
)

// primeira letra maiuscula significa que pode ser usada fora do pacote (analogia public / private)
func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")

	// aqui nao precisa usar prefixo "auxiliar." pois a funcao esta no mesmo modulo
	escrever2()
}
