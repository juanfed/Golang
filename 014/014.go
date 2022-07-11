// repaso del tema de los punteros

// vamos a usar un ejemplo de algo que ya habia echo el cual me hacia una copia de la variable mas no me modificaba la variable como si, devido a que estamos trabajando en bloques diferentes y este realiza una copia y no modifica el valor de la variable original
package main

import "fmt"

// con el * le decimos que ya no va a recibir como argumento una copia si no la variable original
func modificarVariable(parametro *string) { // apuntramos a la variable original
	// ponemos un * para asegurarnos de que estamos asignando un nuevo valor a esa variable original y no a una copia
	*parametro = "Cambio de valor"
}

func main() {
	var curso = "Curso de Golang"
	// con el & le decimos que vamos a pasar ya no una copia si no que le pasamos un espacio en memoria
	fmt.Println("antes de entrar en la fucion\n", curso)
	modificarVariable(&curso)

	fmt.Println(&curso) // referemcia en memoria, ya no pasa como copia

	fmt.Print("despues de la funcion \n", curso)
}
