package main

import "fmt"

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func main() {
	fmt.Printf("%d! = %d \n", 4, factorial(4))
	fmt.Printf("%d! = %d \n", 5, factorial(5))
	fmt.Printf("%d! = %d \n", 7, factorial(7))
}
