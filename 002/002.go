// Tipos de datos y variables en go

package main

import "fmt"

const Euler float32 = 2.71828 // cforma de declarar constantes en go, se hacen con uupercamelcase, fuera de la funcion principal y con la palabra reservada const

func main() {
	fmt.Println("strin")     //string
	fmt.Println(7)           // int
	fmt.Println(7.5)         // esto podria ser un float64 o float32 (raro la verdad :v)
	fmt.Println(false, true) // el booleano de toda la vida

	// definicion de variables
	var texto string
	texto = "texto cualquiera"

	fmt.Println(texto) // si no llamo la variable me marca error XD, no puedo definir variables por definir solamente, tengo que usarlo si o si, no vasta con darle un valor

	var a, b = 3, 5 // declaracion y asignacion de multiples variables
	suma := a + b
	resta := a - b
	fmt.Println(suma)
	fmt.Printf("La resta de %d menos %d es de: %d ", a, b, resta) // %d para números enteros, %s para strings y %f para numeros flotantes, por defecto tambien podemos utilizar %v para todo y el ya por detras asignará el tipo tipo de %signo necesario

	// defirni multiples variables

	var (
		PI    float32 = 3.141592
		radio int     = 5
	)
	// calculo de la longitud de una circunferencia
	fmt.Println("La circunferencia del ciruclo es", 2*PI*float32(radio)) // convierto de entero a flotante, forma de hacerlo aca

}
