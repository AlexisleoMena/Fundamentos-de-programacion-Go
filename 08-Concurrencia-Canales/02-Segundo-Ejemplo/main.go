package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func get(num int) {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/" + strconv.Itoa(num))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	// fmt.Println("status: ", res.Status)

	scanner := bufio.NewScanner(res.Body)
	for i := 0; scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func main() {
	inicio := time.Now()
	for i := 0; i < 100; i++ {
		get(i)
	}
	fmt.Println("Tiempo de ejecución:", time.Since(inicio))

	var terminar string
	fmt.Scan(&terminar)
}
