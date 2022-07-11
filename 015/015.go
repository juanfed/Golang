// funciones recursivas (funciones que se llaman a si mismas) y funciones como valores
package main

import "fmt"

func factorial(number int) int {
	if number == 1 {
		return 1
	}

	return factorial(number-1) * number
}

func factorial2(number int) int {
	if number == 1 {
		return 1
	}

	return factorial(number-1) * number
}

func main() {
	result := factorial(5)

	fmt.Println("el factorial de 5 es:", result)

	// ahora funciones como valores, es decir variables de tipo funcion
	// lo primero es almacenar la funcion en una variable
	var miFuncion = factorial2
	// ahora llamo a a la funcion  factorial2 desde la variable

	result2 := miFuncion(7)
	fmt.Println(result2)
}
