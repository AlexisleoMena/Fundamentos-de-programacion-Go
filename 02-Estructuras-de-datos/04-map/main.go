package main

import "fmt"

func main() {
	//DECLARAR
	dias := make(map[int]string)
	tiempos := make(map[string]float64)
	materias := make(map[string][]string)

	//DEFINIR
	dias[1] = "Lunes"
	dias[7] = "Domingo"

	tiempos["Alexis"] = 33.2
	tiempos["Juan"] = 35.2
	tiempos["Kevin"] = 40.2

	materias["Alexis"] = []string{"Lengua", "Ingles", "Fisica"}
	materias["Pedro"] = []string{"Lengua", "Ingles"}

	fmt.Printf("length: %d, dias: %v \n", len(dias), dias)
	fmt.Printf("length: %d, tiempos: %v \n", len(tiempos), tiempos)
	fmt.Printf("length: %d, materias: %v \n", len(materias), materias)

	//INICIALIZAR
	precios := map[string]float32{
		"Azucar": 55.2,
		"Yerba":  200.50,
		"Sal":    55.2,
	}
	fmt.Printf("length: %d, precios: %v \n", len(precios), precios)

	//IMPRIMIR
	fmt.Println("materias de Alexis:", materias["Alexis"])
	fmt.Println("tiempos de Alexis:", tiempos["Alexis"])
	fmt.Println("precio de Yerba:", precios["Yerba"])

	//ELIMINAR
	delete(precios, "Azucar")
	fmt.Printf("length: %d, precios: %v \n", len(precios), precios)

	delete(tiempos, "Alexis")
	fmt.Printf("length: %d, tiempos: %v \n", len(tiempos), tiempos)

	//BUSCAR
	precio, existe := precios["Yerba"]
	fmt.Printf("existe: %v, precio de Yerba: %v \n", existe, precio)

	precio, existe = precios["Azucar"] //precio = 0, existe = false
	fmt.Printf("existe: %v, precio de Azucar: %v \n", existe, precio)

}
