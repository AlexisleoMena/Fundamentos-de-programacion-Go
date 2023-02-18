package main

import "fmt"

func saludar(nombre string, edad int) {
	fmt.Printf("Hola, soy %s y tengo %d años.\n", nombre, edad)
}

func suma(numA, numB int) int {
	return numA + numB
}

func main() {
	saludar("Alexis", 23)
	fmt.Printf("%d + %d = %d \n", 12, 34, suma(12, 34))
}
