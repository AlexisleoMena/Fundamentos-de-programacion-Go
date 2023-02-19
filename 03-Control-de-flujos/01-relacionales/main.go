package main

import "fmt"

func main() {
	//No esta permitida la comparación de datos de distinto tipo

	var numA, numB int = 4, 5

	fmt.Println("numA == numB", numA == numB)
	fmt.Println("numA != numB", numA != numB)
	fmt.Println("numA < numB", numA < numB)
	fmt.Println("numA > numB", numA > numB)
	fmt.Println("numA <= numB", numA <= numB)
	fmt.Println("numA >= numB", numA >= numB)
}
