package main

import "fmt"

func main() {

	//No existe el loop While

	//Emulación del loop While
	var i int
	for i < 5 {
		fmt.Println("Ejecución de bloque (loop While)")
		i++
	}

	//Estructura general del loop For
	for i := 0; i < 5; i++ {
		fmt.Println("Ejecución de bloque (loop For)")
	}

}
