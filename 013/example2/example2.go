package main

import (
	"fmt"
)

//ejemplo de una fucion con retorno

// el tener (number1, number2 int) será lo mismo que tener (number1 int, number2 int)
// el tipo de valor a retornar se pone luego de los parentesis y la cantidad de tipos de valores a retornar sera igual a la cantidad de valores que retornará la funcion
func suma(number1, number2 int) int {
	// con la palabra return espesicicamos que lo que esta allí es lo que se irá a retornar
	return number1 + number2
}

// una función que me retornará más de un valor
// cuando se retorna mas de un tipo de dato, estos iran al final y dentro de parentesis
func resta(number1 int, number2 int) (int, string) {
	// separo los tipos de datos a retornar por medio de nuna coma ","
	return number1 - number2, "Este es el resultado de la resta"
}

func main() {
	// el orden aca será el mismo orden que como esta en la funcion
	fmt.Println(suma(24, 72))
	// o también puedo asignarle esa funcion a una variable y usar la variable, es lo mas común.
	result := suma(12, 32)
	fmt.Println("La segunda suma es: ", result)

	// tengo que alamacenar el otro tipo de dato y es por eso que declaro otra variable aparte de result2. Si no queremos usar la segunda variable, solo ponemos el guión bajo y listo
	result2, mensaje := resta(24, 13)
	fmt.Print(mensaje, " ", result2)
}
