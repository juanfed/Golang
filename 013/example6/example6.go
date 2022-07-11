// bloques y alcances de las variables
// un bloque solo son una coleccion de sentencias, las cuales estan agrupadas por las llaves
// las variables estan ligadas a los bloques, ya que se crean, modifican y se eliminan dentro del mismo
package main

import "fmt"

func main() {
	// estamos dentro dle bloque main

	var variable1 = "Estoy aprendiendo Golang" // declaracion de varianle dentro del bloque main
	fmt.Println(variable1)                     // llamo la variable dentro del mismo bloque

	// segundo bloque

	{ // como este bloque esta dentro del bloque main, este bloque le pertenece al bloque main y por ende esta imprecion será valida.
		fmt.Println(variable1)

		//tercer bloque
		{ // puedo hacer uso de las variables declaras en los bloques superiores ya que este es un bloque hijo, por ende el alcance de las variables es desendiente hacia los hijos
			var variable3 = "Soy un mensaje de un bloque inferior"
			fmt.Println(variable1)
			fmt.Println(variable3)
		}

		// si hago esto => fmt.Println(variable3) ==> me generará un error ya que fue declara dentro del bloquer 3 y solo puede ser usada dentro de ella o de sus hijos, ya que una vez se termina ese bloque la variable es destruida
		// una vez que un bloque finaliza, todas las variables declaradas en el son destruidas
	}

	// ejemplo de los valores de las variables en los bloques

	// primer bloque
	{
		var bloque = "segundo bloque"

		//segundo bloque
		{
			var bloque = "segundo tercer bloque"

			fmt.Println(bloque) // imprimera el valor que tienne esa variable en este bloque y no el valor que almaceno esa variable en el bloque de arriba
		}

		fmt.Println(bloque)
	}
}
