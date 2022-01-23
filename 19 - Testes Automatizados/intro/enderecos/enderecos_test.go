package enderecos_test

import (
	// . "intro/enderecos" # colocando o "." indica que é o pacote principal, não precisaria colocar "enderecos." antes das funcoes
	"intro/enderecos"
	"testing"
)

func TestCheckTipoEndereco(t *testing.T) {
	fakeEndereco := "Rua teste"
	expected := "Rua"

	actual := enderecos.CheckTipoEndereco(fakeEndereco)

	if actual != expected {
		t.Errorf("CheckTipoEndereco fail. Actual: %s / Expected: %s", actual, expected)
	}
}

type enderecoDataProvider struct {
	enderecoMock string
	expected     string
}

func TestCheckTipoEnderecosComStruct(t *testing.T) {
	mocks := []enderecoDataProvider{
		{"Rua test", "Rua"},
		{"Avenida pra lá", "Avenida"},
		{"Estrada Real", "Estrada"},
		{"Rodovia Real", "Rodovia"},
		{"RUA MAIUSCULA", "Rua"},
		{"NonExiste", "Tipo inválido"},
		{"", "Tipo inválido"},
	}

	for _, mock := range mocks {
		actual := enderecos.CheckTipoEndereco(mock.enderecoMock)

		if actual != mock.expected {
			t.Errorf("CheckTipoEndereco fail. Actual: %s / Expected: %s", actual, mock.expected)
		}
	}
}
