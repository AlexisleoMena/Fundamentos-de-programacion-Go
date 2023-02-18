package main

import "fmt"

func main() {
	hola := "hola"
	mundo := "mundo"

	fmt.Println(hola) //Println realiza saltos de línea
	fmt.Println(mundo)

	fmt.Print(hola, "\n") //Print no realiza saltos de línea
	fmt.Print(mundo, "\n")

	nombre := "Alexis"
	edad := 24

	fmt.Printf("Hola, soy %v y tengo %v años.\n", nombre, edad) //Printf permite imprimir datos formateados en una cadena

	fmt.Printf("Hola, soy %s y tengo %d años.\n", nombre, edad)

	message := fmt.Sprintf("Hola, soy %s y tengo %d años.\n", nombre, edad)
	fmt.Println(message)

	fmt.Printf("nombre: %s,TypeOf: %T \n", nombre, nombre)

	var telefono string
	fmt.Print("Ingrese telefono: ")
	fmt.Scanln(&telefono) //Scanln la entrada de datos o lectura de datos por teclado
	fmt.Println("Su telefono es: ", telefono)

}
