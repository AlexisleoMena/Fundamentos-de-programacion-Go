package main

import "fmt"

func main() {
	var num int // num = 0

	// ++num //No existe pre-incremento
	// --num //No existe pre-decremento

	num++ //num = num + 1
	fmt.Println("num + 1 =", num)

	num-- //num = num - 1
	fmt.Println("num - 1 =", num)

}
