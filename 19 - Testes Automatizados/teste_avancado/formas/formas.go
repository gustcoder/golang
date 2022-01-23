package formas

import (
	"fmt"
	"math"
)

type Forma interface {
	Area() float64
}

func EscreverArea(f Forma, tipo string) {
	fmt.Printf("A área do %s é %0.2f\n", tipo, f.Area())
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

type Circulo struct {
	raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}
