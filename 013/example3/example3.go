// funciones anonimas en golang

package main

import (
	"fmt"
)

func main() {
	// este ipo de funciones no tiene un nombre y van dentro del main

	// ejemplo de una funcion anonima sin parametros
	func() {
		fmt.Println("Hola mundo desde una funcion anonima")
	}() // este parentesis hara el llamado a la funcion

	// otra forma es almacenar la funcion en una variable

	funcion := func() {
		fmt.Println("Función almacenada dentro de una variable")
	}
	// la ventaja es que podemos hacer uso de ella en cualquier parte de nuestro código
	funcion() // de esta forma la llamo

	// ahora los mismo pero con parametros
	miFuncion := func(username string) {
		fmt.Println("Hola", username, "un gusto conocerte")
	}
	// puedo hacer el llamado cuentas veces lo requiera
	miFuncion("Alejandra")
	miFuncion("Alejandra")
	miFuncion("Alejandra")

	// ahora que podamos retornar algo
	miFuncion2 := func(number1, number2 int, user string) (int, string) {
		texto := fmt.Sprintf("Hola %s te saludo desde una funcion sin nombre\n", user)
		// tambien podemos usar el sprintf para dar formato, uso esto y no el println ya que este me devuelve 2 datos y solo necesito uno
		return number1 + number2, texto
	}

	suma, texto := miFuncion2(12, 24, "Daniela")
	fmt.Println(texto, suma)
}
