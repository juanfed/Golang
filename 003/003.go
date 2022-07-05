package main

import (
	"fmt"
	"reflect"
)

// eclaro multiples constantes
const (
	Domingo int = iota // con esto le digo que sera una secuencia, el valor por defecto sera 0 y se irá incrementando en 1., tambien puedo sumarle un numero para empezar desde otro numero, por ejemplo iota + 10 para que inicie en dicho valor
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	fmt.Println(Domingo)
	fmt.Println(Lunes)
	fmt.Println(Martes)
	fmt.Println(Miercoles)
	fmt.Println(Jueves)
	fmt.Println(Viernes)
	fmt.Println(Sabado)

	nombre := "Juan Fernando Gutierrez"

	longitud := len(nombre) // puedo saber el tamaño de los string con la palabra reservada len

	fmt.Println("La longitud del nombres es: ", longitud)

	primerCaracter := nombre[0]
	ultimoCaracter := nombre[len(nombre)-1]

	fmt.Printf("El primer caracter es: %d \n", primerCaracter) // me imprime un número que hace referncia a ese caracter en la tabla ACII
	fmt.Println(reflect.TypeOf(primerCaracter))                // entero positivo de obits, un caracter como tal

	// forma para hacer que si o si me muestre el caracter
	fmt.Printf("El primer caracter es: %c \n", primerCaracter)
	fmt.Printf("El ultimo caracter es: %c \n", ultimoCaracter)
}
