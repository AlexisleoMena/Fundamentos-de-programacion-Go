package main

import "fmt"

func main() {

	//DECLARAR
	var edad int       //0
	var nombre string  //""
	var altura float64 //0
	var esSoltero bool //false
	//DEFINIR
	edad = 24
	nombre = "Alexis"
	altura = 1.75
	esSoltero = true

	//INICIALIZAR
	var apellido string = "Mena"
	telefono := 3424061761 //Infiere que es int

	const PI = 3.1415

	fmt.Println(nombre, apellido, edad, altura, esSoltero, telefono, PI)

}
