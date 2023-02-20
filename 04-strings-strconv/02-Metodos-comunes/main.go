package main

import (
	"fmt"
	"strconv"
)

func main() {

	//STRCONV
	var nombre, edad, altura, casado string = "Alexis", "22", "1.74", "false"

	fmt.Printf("%T, %v\n", nombre, nombre)

	if s, err := strconv.Atoi(edad); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

	if s, err := strconv.ParseInt(edad, 10, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

	if s, err := strconv.ParseFloat(altura, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

	if s, err := strconv.ParseBool(casado); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
}
