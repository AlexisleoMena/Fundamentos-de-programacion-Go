package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("El servidor esta corriendo en el puerto 3001")
	fmt.Println("http://localhost:3001")
	log.Fatal(http.ListenAndServe("localhost:3001", nil))

}
