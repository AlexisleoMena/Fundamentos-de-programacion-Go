package main

import "fmt"

func main() {
	var edad int = 19

	// if edad >= 18 fmt.Println("Es mayor de edad") //No es posible obviar las llaves

	if edad >= 18 {
		fmt.Println("Es mayor de edad")
	} else {
		fmt.Println("No es mayor de edad")
	}

	var num int = 12
	if num%4 == 0 {
		fmt.Printf("El numero %d es divisible por 4\n", num)
	}

	var passA string = "1234"
	if passB := "1234"; passA == passB {
		fmt.Printf("Las contraseñas %s y %s coinciden\n", passA, passB)
	} else {
		fmt.Printf("Las contraseñas %s y %s coinciden\n", passA, passB)
	}

	var usuarioA, pass string = "Alexis", "1234"
	if usuarioB := "Alexis"; usuarioA == usuarioB {
		if passB := "1234"; pass == passB {
			fmt.Printf("usuario %s y contraseña %s coinciden", usuarioA, pass)
		}
	}
}
