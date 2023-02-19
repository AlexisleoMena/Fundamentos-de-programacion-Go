package main

import "fmt"

func main() {
	var num int = 2

	num += 2
	fmt.Println("num + 2 = ", num)

	num -= 2
	fmt.Println("num - 2 = ", num)

	num /= 2
	fmt.Println("num / 2 = ", num)

	num *= 2
	fmt.Println("num * 2 = ", num)

	num %= 2
	fmt.Println("num % 2 = ", num)

}
