// variables globales
// es decir como multiples funciones pueden acceder a la misma variable

package main

import "fmt"

// voy a crear unas funciones que accederan a la misma variable, podrán leerla y modificarla

var username string // al ser declarada con var y fuera de cualquier funcion, se crea como una variable de alcance global

func funcion1() {
	username = "nombre 1"
	fmt.Println("Funcion 1", username)
}

func funcion2() {
	username = "nombre 2"
	fmt.Println("Funcion 2", username)
}
func main() {
	// aca modifico el valor de la variable
	username = "Juan"

	// ambas funciones tienen acceso a la variable
	funcion1()
	funcion2()

}
