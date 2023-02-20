package main

import (
	"fmt"
	"os"
)

func readFile(pathFile string) string {
	if file, err := os.Open(pathFile); err != nil {
		panic("No fue posible obtener el archivo!") //"panic" aborta el programa.
	} else {
		defer func() {
			fmt.Println("Cerrando el archivo!")
		}()
		contenido := make([]byte, 256)
		long, _ := file.Read(contenido)
		return string(contenido[:long])
	}
}

func recoverFromPanic() {
	if err := recover(); err != nil { //"recover" recupera el programa de un bloqueo por ejecución de "panic". Evita que "panic" aborte el programa.
		fmt.Println("Recuperado de error:", err)
	}
}

func survivingThePanic(pathName string) {
	defer recoverFromPanic()
	readFile(pathName)
	//Lo que escriba acá no se ejecutará si el path es incorrecto.
}

func main() {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("Recuperado de error:", err)
	// 	}
	// }()
	// readFile("nombress.txt")
	// fmt.Println("Si hay pánico no sobrevivo!") //No se ejecuta si el path es incorrecto.

	survivingThePanic("nombres.txt")
	fmt.Println("Sobreviví al pánico!") //Se ejecuta aunque el path sea incorrecto.

}
