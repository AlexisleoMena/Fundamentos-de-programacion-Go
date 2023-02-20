package main

import (
	"fmt"
	"os"
)

func main() {
	if file, err := os.Open("hola.txt"); err != nil {
		fmt.Println("No fue posible obtener el archivo!")
	} else {
		contenido := make([]byte, 254)  //El contenido se guardará en un array de 254 elementos de tipo byte.
		long, _ := file.Read(contenido) //long = cantidad de carácteres. _ = err. contenido = [72 111 108 97 32 97 32 116 111 100 111 115 33, ...0]
		str := string(contenido[:long]) //Es necesario parsear el array a string.
		fmt.Println(str)
	}
}
