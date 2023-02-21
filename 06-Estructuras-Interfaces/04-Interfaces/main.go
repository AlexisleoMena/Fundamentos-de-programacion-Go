package main

import "fmt"

type Animal interface { //Perro, Gato y Persona implementan la interfaz Animal solo sí implementan todos los metodos declarados en la firma.
	Sonido() string
}

type Perro struct{}
type Gato struct{}
type Persona struct{}

func (*Perro) Sonido() string {
	return "¡Guau Guau!"
}
func (*Gato) Sonido() string {
	return "¡Miau Miau!"
}
func (*Persona) Sonido() string {
	return "¡Hola!"
}

func emitirSonido(animal Animal) { //Puede recibir tipos que implementen la interfaz Animal.
	fmt.Println(animal.Sonido())
}

//--------------------------------

type Figura interface { //Cuadrado y Circulo implementan la interfaz Figura solo sí implementan todos los metodos declarados en la firma.
	ObtenerArea() float64
	ObtenerPerimetro() float64
}

type Cuadrado struct {
	Lado float64
}
type Circulo struct {
	Radio float64
}

func (c *Cuadrado) ObtenerArea() float64 {
	return c.Lado * c.Lado
}
func (c *Cuadrado) ObtenerPerimetro() float64 {
	return 4 * c.Lado
}
func (c *Circulo) ObtenerArea() float64 {
	return 3.1415 * c.Radio * c.Radio
}
func (c *Circulo) ObtenerPerimetro() float64 {
	return 2 * 3.1415 * c.Radio
}

func GetAreaYPerimetro(figura Figura) { //Puede recibir tipos que implementen la interfaz Figura.
	fmt.Printf("\nArea: %v, Perimetro: %v \n", figura.ObtenerArea(), figura.ObtenerPerimetro())
}

func main() {
	perro1 := Perro{}
	gato1 := Gato{}
	persona1 := Persona{}

	emitirSonido(&perro1)
	emitirSonido(&gato1)
	emitirSonido(&persona1)

	// --------------------------

	cuadrado1 := Cuadrado{5}
	circulo1 := Circulo{5}

	GetAreaYPerimetro(&cuadrado1)
	GetAreaYPerimetro(&circulo1)
}
