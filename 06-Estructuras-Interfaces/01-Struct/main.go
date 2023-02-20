package main

import "fmt"

//"Go es un lenguaje no orientado a objetos que permite la programación orientada a objetos - aunque no de la forma tradicional"
//Pilares de la POO
// - Abstracción: No hay clases, hay struct.
// - Encapsulamiento: No hay etiquetas "public", "private" o "protected". Se emula mediante iniciales mayúsculas o minúsculas.
// - Herencia: No hay herencia, hay "composición".
// - Polimorfismo: Mediante Interfaces.

type Persona struct {
	nombre string
	edad   int
}

func (p *Persona) GetNombre() string {
	return p.nombre
}

func (p *Persona) SetNombre(nombre string) {
	p.nombre = nombre
}

func main() {
	// var p1 Persona //p1 = {"", 0}
	// p1 = Persona{"Alexis", 24}
	p1 := Persona{"Alexis", 24}
	fmt.Println("p1.getNombre() =", p1.GetNombre())

	p2 := Persona{
		nombre: "Juan",
		edad:   24,
	}
	p2.SetNombre("Pedro")
	fmt.Println("p2.getNombre() =", p2.GetNombre())
}
