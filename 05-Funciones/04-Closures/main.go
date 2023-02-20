package main

import (
	"fmt"
	"math"
)

func saludar(yo string) func(persona string) {
	return func(persona string) {
		fmt.Printf("Hola %s, me llamo %s\n", persona, yo)
	}
}

func potencia(base float64) func(exponente float64) float64 {
	return func(exponente float64) float64 {
		return math.Pow(base, exponente)
	}
}

func main() {
	AlexisSaludo := saludar("Alexis")
	AlexisSaludo("Antonela")
	AlexisSaludo("Agustin")

	potenciaEnBaseDos := potencia(2)
	fmt.Printf("2^%v = %v \n", 1, potenciaEnBaseDos(1))
	fmt.Printf("2^%v = %v \n", 3, potenciaEnBaseDos(3))

	potenciaEnBaseTres := potencia(3)
	fmt.Printf("3^%v = %v \n", 1, potenciaEnBaseTres(1))
	fmt.Printf("3^%v = %v \n", 3, potenciaEnBaseTres(3))
}
