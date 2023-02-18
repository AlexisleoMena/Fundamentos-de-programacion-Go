package main

import (
	"fmt"
	"reflect"
)

func main() {
	var numA int = 9
	var numB int = 4

	suma := numA + numB
	diferencia := numA - numB
	producto := numA * numB

	fmt.Println("suma:", suma, diferencia, "TypeOf:", reflect.TypeOf(suma))
	fmt.Println("diferencia:", diferencia, "TypeOf:", reflect.TypeOf(diferencia))
	fmt.Println("producto:", producto, "TypeOf:", reflect.TypeOf(producto))

	cocienteA := numA / numB //Infiere que cocienteA es de tipo int
	fmt.Println("cocienteA:", cocienteA, "TypeOf:", reflect.TypeOf(cocienteA))

	var numC float64 = 7.0
	var numD float64 = 2.0

	cocienteB := numC / numD
	cocienteC := numA / int(numD)     //Infiere que cocienteC es de tipo int
	cocienteD := float64(numA) / numD //Infiere que cocienteD es de tipo float64

	fmt.Println("cocienteB:", cocienteB, "TypeOf:", reflect.TypeOf(cocienteB))
	fmt.Println("cocienteC:", cocienteC, "TypeOf:", reflect.TypeOf(cocienteC))
	fmt.Println("cocienteD:", cocienteD, "TypeOf:", reflect.TypeOf(cocienteD))

	// cocienteE := numA / numC //Operación invalida al ser variables de distinto tipo
	cocienteE := 12 / 2.5 //Infiere que cocienteE es de tipo float64
	fmt.Println("cocienteE:", cocienteE, "TypeOf:", reflect.TypeOf(cocienteE))

	modulo := 12 % 5
	fmt.Println("modulo:", modulo, "TypeOf:", reflect.TypeOf(modulo))

}
