package main

import "fmt"

func main() {
	var letra string = "k"

	switch {
	case letra == "a" || letra == "A": //Este tipo de switch-case permite cases de tipo boolean
		fmt.Printf("La letra %s es vocal \n", letra)
	case letra == "e" || letra == "E":
		fmt.Printf("La letra %s es vocal \n", letra)
	case letra == "i" || letra == "I":
		fmt.Printf("La letra %s es vocal \n", letra)
	case letra == "o" || letra == "O":
		fmt.Printf("La letra %s es vocal \n", letra)
	case letra == "u" || letra == "U":
		fmt.Printf("La letra %s es vocal \n", letra)
	default:
		fmt.Printf("La letra %s no es vocal \n", letra)
	}

	//Este tipo de switch-case permite cases de igual tipo al del parametro que se indica al comienzo
	switch letra {
	case "a", "A", "e", "E", "i", "I", "o", "O", "u", "U": //como "letra" es un string =>  cada case solo puede aceptar strings
		fmt.Printf("La letra %s es vocal \n", letra)
	default:
		fmt.Printf("La letra %s no es vocal \n", letra)
	}

	var num int = 4
	switch {
	case num%2 == 0:
		fmt.Printf("El numero %d es par \n", num) //Auto-break
	case num > 0:
		fmt.Printf("El numero %d es positivo \n", num)
	default:
		fmt.Printf(":( no se dio ningún caso \n")
	}
}
