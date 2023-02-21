package mensajes

import "fmt"

//Los métodos y variables con inicial minúscula son privadas, no exportables fuera del paquete "mensajes".

func metodoPrivado() {
	fmt.Println("Metodo no exportable, se utiliza solo en el paquete mensajes.")
}

var nombre string = "Alexis"
var edad int = 24

func Saludar() {
	fmt.Println("Hola a todos!")
}

func QuienSoy() {
	fmt.Printf("\nMe llamo %s y tengo %d años.", nombre, edad)
}
