// repaso condicionales (controles de flujo)
package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("Ingrese su nombre: ")
	fmt.Scanf("%s\n", &name) // si no pongo el salto de linea no me funciona para leer la edad
	fmt.Print("Ingrese su edad: ")
	fmt.Scanf("%d\n", &age)

	if age >= 18 {
		fmt.Println("Eres mayor de edad, puedes pasar 😎")
	} else {
		fmt.Println("No eres mayor de edad, no puedes pasar 😥")
	}

}
