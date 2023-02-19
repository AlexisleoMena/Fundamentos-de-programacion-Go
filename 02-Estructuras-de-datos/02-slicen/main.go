package main

import "fmt"

func main() {

	//DECLARAR
	var nombres []string
	var edades []int

	//DEFINIR
	nombres = []string{"Alexis", "Juan"}
	edades = []int{22, 32}
	fmt.Println("nombres:", nombres, ", length:", len((nombres)))
	fmt.Println("edades:", edades, ", length:", len((edades)))

	//INICIALIZAR
	var colores []string = []string{"Azul", "Rojo"}
	frutas := []string{"Manzana", "Naranja", "Pera"}
	fmt.Println("colores:", colores, ", length:", len((colores)))
	fmt.Println("frutas:", frutas, ", length:", len((frutas)))

	//PUSHEAR/APENDEAR
	paises := []string{"Perú", "Chile", "Uruguay"}
	paises = append(paises, "Brazil")
	paises = append(paises, "Argentina", "Colombia")
	fmt.Println("paises:", paises, ", length:", len((paises)))

	//SUB-SLICEN Y REFERENCIAS
	provincias := []string{"Santa Fe", "Corritentes", "Cordoba", "Jujuy"}
	subSlicenProvincias := provincias[1:3] //["Corritentes", "Cordoba"]
	provincias[2] = "Salta"                // provincias[2] = "Cordoba" y subSlicenProvincias[1] = "Cordoba" => provincias[2] = "Salta" y subSlicenProvincias[1] = "Salta"
	fmt.Println("provincias:", provincias)
	fmt.Println("subSlicenProvincias:", subSlicenProvincias)

	//PUNTERO, LONGITUD Y CAPACIDAD

	meses := []string{"Enero", "Febrero", "Marzo"}

	fmt.Printf("Referencia / dirección de memoria al dato del primer elemento del segmento: %p \n", meses)
	fmt.Printf("Longitud / cantidad de elementos que posee el segmento: %v \n", len(meses))
	fmt.Printf("Capacidad / cantidad de elementos que el segmento puede poseer: %v \n\n", cap(meses))

	meses = append(meses, "Abril")
	fmt.Printf("Referencia / dirección de memoria al dato del primer elemento del segmento: %p \n", meses)
	fmt.Printf("Longitud / cantidad de elementos que posee el segmento: %v \n", len(meses))
	fmt.Printf("Capacidad / cantidad de elementos que el segmento puede poseer: %v \n", cap(meses))
	//Si al apendear la nueva longitud supera la capacidad del segmento, entonces la función append genera
	//un nuevo slicen con otra referencia y mayor capacidad.

}
