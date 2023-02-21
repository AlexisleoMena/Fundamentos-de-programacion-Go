package figuras

import "fmt"

type Figura interface {
	obtenerArea() float64
	obtenerPerimetro() float64
}

func GetAreaYPerimetro(figura Figura) {
	fmt.Printf("\nArea: %v, Perimetro: %v \n", figura.obtenerArea(), figura.obtenerPerimetro())
}
