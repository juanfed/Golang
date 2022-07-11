// funciones como argumentos
package main

import "fmt"

// en go nosotros podemos tratar las funciones como argumentos o reternar otras funciones

// para las funciones como argumentos necesitaremos crear un nuevo tipo de dato

// con esto le digo a go que toda funcion que reciba 2 parametros como enteros y que retorne un entero sera catalogada como una funcion de tipo operacion
type Operacion func(balance, cantidad int) int

// cumple con los requisitos de operación, recibe 2 enteros y retorna un entero también

func resta(number1, number2 int) int { // catalogada como una funcion de tipo operacion
	return number1 - number2
}

func suma(number3, number4 int) int { // catalogada como una funcion de tipo operacion
	return number3 + number4
}

// (funcion Operacion) => defino que la palabra funcion sera un parametro de tipo Operacion sea ya suma o resta
func ejecutaroOperacion(funcion Operacion, balance, cantidad int) {
	fmt.Println("Antes de la operacion")

	result := funcion(balance, cantidad)

	fmt.Println("El resultado de la operacion es: ", result)

	fmt.Println("Después de la operación")
}

func main() {
	ejecutaroOperacion(suma, 14, 24)
	ejecutaroOperacion(resta, 14, 24)
}
