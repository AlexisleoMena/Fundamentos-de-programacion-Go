package main

import (
	"fmt"
	"paquetes/figuras"
	"paquetes/mensajes"
	"paquetes/models"
)

func main() {
	mensajes.Saludar()
	mensajes.QuienSoy()

	circ1 := figuras.Circulo{Radio: 3}
	cuad1 := figuras.Cuadrado{Lado: 2}

	figuras.GetAreaYPerimetro(&circ1)
	figuras.GetAreaYPerimetro(&cuad1)

	p1 := models.Persona{}
	p1.Constructor("Alexis", 23)
	fmt.Println(p1.GetNombre())
	fmt.Println(p1.GetEdad())

}
