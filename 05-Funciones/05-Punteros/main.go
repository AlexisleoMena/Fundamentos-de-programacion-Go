package main

import "fmt"

func modificarParametro(pointer *string, nombre string) {
	*pointer = nombre
}

func main() {
	var nombre string = "Alexis"
	var pointer *string = &nombre

	fmt.Println("nombre:", nombre)
	fmt.Println("&nombre:", &nombre)
	fmt.Println("*pointer:", *pointer)
	fmt.Println("pointer:", pointer)

	*pointer = "Juan"

	fmt.Println()
	fmt.Println("nombre:", nombre)
	fmt.Println("&nombre:", &nombre)
	fmt.Println("*pointer:", *pointer)
	fmt.Println("pointer:", pointer)

	modificarParametro(&nombre, "José")

	fmt.Println()
	fmt.Println("nombre:", nombre)
	fmt.Println("&nombre:", &nombre)
	fmt.Println("*pointer:", *pointer)
	fmt.Println("pointer:", pointer)

	modificarParametro(pointer, "Esteban")

	fmt.Println()
	fmt.Println("nombre:", nombre)
	fmt.Println("&nombre:", &nombre)
	fmt.Println("*pointer:", *pointer)
	fmt.Println("pointer:", pointer)

}
