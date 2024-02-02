package mypackage

type Figura interface {
	Area() float64
}

type Cuadrado struct {
	Base float64
}

type Triangulo struct {
	Base   float64
	Altura float64
}

func (c Cuadrado) Area() float64 {
	return c.Base * c.Base
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) / 2
}

func Calcular(f Figura) float64 {
	return f.Area()
}
