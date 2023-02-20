package main

import "fmt"

func main() {

	//USO DE BREAK
	nombres := [...]string{
		"Alexis",
		"Juan",
		"Carlos",
		"Esteban",
	}
	nombre := "Carlos"
	var indice int = -1
	for i := 0; i < len(nombres); i++ {
		fmt.Println("Inicio de la iteración", i+1) //Hará 3 iteraciones hasta encontrar a Carlos, luego sale del loop.
		if nombres[i] == nombre {
			indice = i
			break
		}
		fmt.Println("Fin de la iteración", i+1)

	}
	if indice != -1 {
		fmt.Printf("El nombre Carlos se encuntra en el indice %d \n", indice)
	} else {
		fmt.Println("El nombre Carlos no se encuentra dentro de la lista de nombres")
	}

	fmt.Println()

	//USO DE CONTINUE
	for i := 1; i < 6; i++ {
		fmt.Println("Inicio de la iteración", i)
		if i%2 == 0 {
			continue //Cada que i sea par, "continue" evita las líneas de código posteriores del loop e inicia la siguiente iteración.
		}
		fmt.Println("Fin de la iteración", i)
	}

}
