package main

import (
	"fmt"
	"net/http"
	"time"
)

func revisarServidor(servidor string) {
	if _, err := http.Get(servidor); err != nil {
		fmt.Println(servidor, "no esta  funcionando")
		return
	}
	fmt.Println(servidor, "esta  funcionando")
}

func main() {
	inicio := time.Now()

	servidores := []string{
		"https://oregoom.com/",
		"https://www.udemy.com/",
		"https://www.youtDDDube.com/",
		"https://www.facebook.com/",
		"https://www.google.com/",
	}

	for _, servidor := range servidores {
		revisarServidor(servidor)
	}

	fmt.Println("Tiempo de ejecución:", time.Since(inicio))
}
