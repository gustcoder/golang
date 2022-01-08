package enderecos

import (
	"strings"
)

// CheckTipoEndereco verifica tipo de endereco
func CheckTipoEndereco(endereco string) string {
	tiposValidos := []string{"Rua", "Avenida", "Estradas", "Rodovia"}

	tipoEndereco := strings.Split(endereco, " ")[0]

	tipoValido := false
	for _, tipo := range tiposValidos {
		if strings.EqualFold(tipo, tipoEndereco) {
			tipoValido = true
		}
	}

	if tipoValido {
		return strings.Title(tipoEndereco) // ucfirst
	}

	return "Tipo inválido"
}
