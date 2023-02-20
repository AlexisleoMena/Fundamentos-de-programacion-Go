package main

import "fmt"

func main() {
	nombres := []string{
		"Alexis",
		"Pedro",
		"Juan",
		"Carlos",
	}

	for i, nombre := range nombres {
		fmt.Printf("nombres[%d] = %s \n", i, nombre)
	}

	for _, nombre := range nombres {
		fmt.Println("nombre: ", nombre)
	}

	for i := range nombres {
		fmt.Println("indice: ", i)
	}
}
