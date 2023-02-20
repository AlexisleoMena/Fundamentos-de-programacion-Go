package main

import (
	"errors"
	"fmt"
)

func cociente(dividendo, divisor int) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("En los números reales no es posible dividir por cero.")
	} else {
		return float64(dividendo) / float64(divisor), nil
	}
}

func main() {
	cociente, err := cociente(3, 2)
	if err == nil {
		fmt.Println("3/2 = ", cociente)
	} else {
		fmt.Println(err)
	}
}
