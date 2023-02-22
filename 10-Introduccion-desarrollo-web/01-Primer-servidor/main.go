package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handlers
func Hola(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Hola mundo!")
}

func NotFound(rw http.ResponseWriter, r *http.Request) {
	http.NotFound(rw, r)
}

func main() {

	//Router
	http.HandleFunc("/", Hola)
	http.HandleFunc("/page", NotFound)

	//Crear servidor
	fmt.Println("El servidor esta corriendo en el puerto 3001")
	fmt.Println("http://localhost:3001")
	log.Fatal(http.ListenAndServe("localhost:3001", nil))

}
