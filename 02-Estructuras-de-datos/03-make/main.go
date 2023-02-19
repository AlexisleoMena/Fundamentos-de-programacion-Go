package main

import "fmt"

func main() {
	meses := make([]string, 0, 3) //make(tipo, longitud, capacidad)
	fmt.Printf("referencia: %p, longitud: %d, capacidad: %d, meses: %v\n", meses, len(meses), cap(meses), meses)

	meses = append(meses, "Enero", "Febrero")
	fmt.Printf("referencia: %p, longitud: %d, capacidad: %d, meses: %v\n", meses, len(meses), cap(meses), meses)

	meses = append(meses, "Marzo", "Abril")
	fmt.Printf("referencia: %p, longitud: %d, capacidad: %d, meses: %v\n", meses, len(meses), cap(meses), meses)

	meses = append(meses, "Mayo", "Junio")
	fmt.Printf("referencia: %p, longitud: %d, capacidad: %d, meses: %v\n", meses, len(meses), cap(meses), meses)

	meses = append(meses, "Julio", "Agosto")
	fmt.Printf("referencia: %p, longitud: %d, capacidad: %d, meses: %v\n", meses, len(meses), cap(meses), meses)
}
