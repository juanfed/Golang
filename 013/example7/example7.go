// variables y scopes en golang
package main

import "fmt"

func modificarVariable(parametro string) { // nos servira para modifgicar el valor de una variable
	fmt.Println("Valor actual: ", parametro)
	parametro = "Cambio de valor"
	fmt.Println("Nuevo valor: ", parametro)
	fmt.Printf("%v\n", &parametro)
}

func main() {
	var curso = "Curso de Golang"
	modificarVariable(curso)
	// se mantendra el valor original y no el valor que queda despues de ejecutarse la funcion
	fmt.Println(curso)
	// esto sucedes porque curso no es la misma variable que parametros, lo podemos comprobar imprimiento sus direcciones de memoria. Porque estamos trabajando con 2 bloques diferentes y al pasar la varibale como argumento una variabnle a una funcion esta pasará como copia y no como referencia en memoria.
	fmt.Printf("%v\n", &curso)
}
