package main

import "fmt"

var mensaje string //"mensaje" puede accederse desde cualquier parte del programa debajo de esta línea.

func fn1() {
	mensaje = "Hola desde fn1!"
	fmt.Println(mensaje)
}
func fn2() {
	mensaje = "Hola desde fn2!"
	fmt.Println(mensaje)
}
func main() {
	mensaje = "Hola desde main!"
	fmt.Println(mensaje)

	defer fn1() //fn1 se ejecuta justo antes de finalizar main
	fn2()

	fmt.Println(mensaje)
}
