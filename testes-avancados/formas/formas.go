package formas

import (
	"math"
)

//Forma
type Forma interface {
	area() float64
}

//Retangulo
type Retangulo struct {
	Altura  float64
	Largura float64
}

//Area
func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

//Circulo
type Circulo struct {
	raio float64
}

//Area
func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}
