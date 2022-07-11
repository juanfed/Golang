// repaso a los retornos de funciones

// aca sera de funciones que retornaran funciones
package main

import "fmt"

// Las funciones que cumplan con este requisito seran de tipo Operacion
type Operacion func(balance, cantidad int) int

// defino las funciones
func crearOperacion(tipo string) Operacion {
	if tipo == "suma" {
		return func(balance1, cantidad1 int) int { return balance1 + cantidad1 }
	} else if tipo == "resta" {
		return func(balance2, cantidad2 int) int { return balance2 - cantidad2 }
	} else {
		return func(balance3, cantidad3 int) int { return balance3 * cantidad3 }
	}
}

func main() {
	suma := crearOperacion("suma")
	resta := crearOperacion("resta")
	multiplicacion := crearOperacion("multiplicacion")

	fmt.Println(suma(10, 20))
	fmt.Println(resta(10, 20))
	fmt.Println(multiplicacion(10, 20))

	// otra forma un poco mas larga seria:
	result := suma(14, 20)
	fmt.Println("El resultadod de la suma es: ", result)

	// esto es util a la hora de implementar patrones de diseño o para tener el código muy bien organizado

}
