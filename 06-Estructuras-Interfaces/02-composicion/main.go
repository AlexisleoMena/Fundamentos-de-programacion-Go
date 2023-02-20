package main

import "fmt"

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
func (p *Persona) GetEdad() int {
	return p.edad
}
func (p *Persona) SetEdad(edad int) {
	p.edad = edad
}

type Alumno struct {
	grado int
	Persona
}

func (a *Alumno) GetGrado() int {
	return a.grado
}
func (a *Alumno) SetGrado(grado int) {
	a.grado = grado
}

func main() {
	alumno1 := Alumno{
		grado:   8,
		Persona: Persona{"Alexis", 24},
	}
	fmt.Printf("Nombre: %s, Edad: %d, Grado: %d\n", alumno1.GetNombre(), alumno1.GetEdad(), alumno1.GetGrado())
}
