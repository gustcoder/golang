package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	// t.Run serve para separar as funcoes que estao sendo testadas
	t.Run("test Retangulo", func(t *testing.T) {
		ret := Retangulo{10, 20}
		actual := ret.Area()

		expected := 200.00

		if actual != expected {
			t.Errorf("Test failed. Actual: %f Expected: %f", actual, expected)
		}
	})

	t.Run("test Circulo", func(t *testing.T) {
		circ := Circulo{10}
		actual := circ.Area()

		expected := float64(math.Pi * 100)

		if actual != expected {
			t.Errorf("Test failed. Actual: %f Expected: %f", actual, expected)
		}
	})
}
