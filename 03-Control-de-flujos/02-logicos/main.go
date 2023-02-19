package main

import "fmt"

func main() {
	var numA, numB, numC int = 1, 2, 3

	fmt.Println("numA > numB && numB > numC", numA > numB && numB > numC)       //false && false => false
	fmt.Println("!(numA > numB) && numB > numC", !(numA > numB) && numB > numC) //true && false => false
	fmt.Println("numC > numA && numC > numB", numC > numA && numC > numB)       //true && true => true

	fmt.Println("numA > numB || numB > numC", numA > numB || numB > numC)       //false || false => false
	fmt.Println("!(numA > numB) || numB > numC", !(numA > numB) || numB > numC) //true || false => true
	fmt.Println("numC > numA || numC > numB", numC > numA || numC > numB)       //true || true => true

}
