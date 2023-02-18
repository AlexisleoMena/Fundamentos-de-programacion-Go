package main

import (
	"fmt"
	"reflect"
)

func main() {

	//DECLARAR
	var edades [3]int       //[0, 0, 0]
	var nombres [2]string   //["", "", ""]
	var alturas [4]float64  //[0, 0, 0]
	var condiciones [2]bool //[false, false, false]

	//DEFINIR
	edades[0] = 33
	edades[1] = 22
	edades[2] = 13
	// edades[3] = 12 //Error por indice invalido
	fmt.Println("edades:", edades, ", length:", len(edades))

	// nombres[0] = ""
	nombres[1] = "Alexis"
	fmt.Println("nombres:", nombres, ", length:", len(nombres))

	// alturas[0] = 0
	alturas[1] = 3.1
	// alturas[2] = 0
	alturas[3] = 4.4
	fmt.Println("alturas:", alturas, ", length:", len(alturas))

	condiciones[0] = true
	condiciones[1] = false
	fmt.Println("condiciones:", condiciones, ", length:", len(condiciones))

	//INICIALIZAR
	var colores [3]string = [3]string{"Azul", "Rojo", "Verde"}
	fmt.Println("colores:", colores, ", length:", len(colores), ", TypeOf:", reflect.TypeOf(colores))

	monedas := [3]string{"Euro", "Peso", "Dolar"} //Infiere que monedas es de tipo array con una longitud de 3 elementos ([3]string)
	fmt.Println("monedas:", monedas, ", length:", len(monedas))

	frutas := [5]string{"Manzana", "Naranja"} // ["Manzana", "Naranja", "", "", ""]
	fmt.Println("frutas:", frutas, ", length:", len(frutas))

	verduras := [4]string{
		0: "Lechuga",
		3: "Zanahoria",
	}
	fmt.Println("verduras:", verduras, ", length:", len(verduras)) // ["Lechuga", "", "", "Zanahoria"]

	marcas := [...]string{"Apple", "Samsung"}
	fmt.Println("marcas:", marcas, ", length:", len(marcas))

	pokemones := [...]string{
		0: "Pikachu",
		2: "Pichu",
		5: "Raichu",
	}
	fmt.Println("pokemones:", pokemones, ", length:", len(pokemones)) // ["Pikachu", "", "Pichu", "", "", "Raichu"]

	paises := [4]string{"Chile", "Perú", "Mexico", "Brazil"}
	subArrayPaisesA := paises[1:3] // Toma el intervalo de indices [1, 3)
	fmt.Println("subArrayPaisesA:", subArrayPaisesA, ", length:", len(subArrayPaisesA))

	subArrayPaisesB := paises[:3] // Toma el intervalo de indices [0, 3)
	fmt.Println("subArrayPaisesB:", subArrayPaisesB, ", length:", len(subArrayPaisesB))

	subArrayPaisesC := paises[1:] // Toma el intervalo de indices [1, len(paises))
	fmt.Println("subArrayPaisesC:", subArrayPaisesC, ", length:", len(subArrayPaisesC))

}
