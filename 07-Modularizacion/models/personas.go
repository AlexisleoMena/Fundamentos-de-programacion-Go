package models

type Persona struct {
	nombre string
	edad   int
}

func (p *Persona) Constructor(nombre string, edad int) {
	p.nombre = nombre
	p.edad = edad
}
