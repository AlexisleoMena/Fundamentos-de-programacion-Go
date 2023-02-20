package main

import (
	"fmt"
	"strings"
)

func saludar(yo string, personas ...string) {
	fmt.Printf("Hola %s, soy %s. \n", strings.Join(personas, ", "), yo)
}

func mejoresTiempos(tiempos ...float64) (float64, float64) {
	var prim, seg float64 = tiempos[0], tiempos[1]
	if prim > seg {
		aux := prim
		prim = seg
		seg = aux
	}
	for i := 2; i < len(tiempos); i++ {
		if prim > tiempos[i] {
			seg = prim
			prim = tiempos[i]
		} else {
			if seg > tiempos[i] {
				seg = tiempos[i]
			}
		}
	}
	return prim, seg
}

func main() {
	saludar("Alexis", "Juan", "Kevin", "Carlos")

	prim, seg := mejoresTiempos(1.33, 2.43, 3.22, 1.12, 1.32)

	fmt.Printf("Mejor tiempo: %v \nSegundo: %v", prim, seg)
}
