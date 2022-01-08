package enderecos

import (
	"testing"
)

func TestCheckTipoEndereco(t *testing.T) {
	fakeEndereco := "Rua teste"
	expected := "Rua"

	actual := CheckTipoEndereco(fakeEndereco)

	if actual != expected {
		t.Errorf("CheckTipoEndereco fail. Actual: %s / Expected: %s", actual, expected)
	}
}
