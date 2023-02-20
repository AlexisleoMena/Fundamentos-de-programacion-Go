package main

import "fmt"

func main() {

	//IIFE
	(func() {
		fmt.Println("Ejecutando IIFE")
	})()

	//Function Expression
	saludar := func() {
		fmt.Println("Hola a todos!")
	}
	saludar()

}
