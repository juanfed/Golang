package main

import "fmt"

func main() {
	// emepzaremos a trabajar con valores de entrada

	// pequeño programa que pida la altura:

	var nombre string
	var edad int8
	var altura float32

	fmt.Print("Ingrese tu nombre: ")
	fmt.Scanf("%s", &nombre)

	fmt.Print("Ingrese tu edad: ")
	fmt.Scanf("%d", &edad)

	fmt.Print("Ingrese su altura: ")
	fmt.Scanf("%.2f", &altura)

}
