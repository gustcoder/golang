package auxiliar

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func escrever2() {
	fmt.Println("Escrevendo pela funcao escrever2")

	erro := checkmail.ValidateFormat("teste@email.com")
	fmt.Println(erro)
}
