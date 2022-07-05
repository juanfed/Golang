/* Foma de hacer un comentario en go en forma de bloque */
// comentario de una linea
// go Es un lenguaje de programacion fuertemente tipado

// forma de hacer una impreción en pantalla.
package main // definimos el paquete (main indica que es el archivo principal de nuestra aplicación)

import "fmt" // paquete para mostrar datos por pantalla, las importaciones se hacen entre comillas.

func main() { // funcion main (funcion principal)
	// acá ira todo lo que se ejecutará

	// var tiene un alcanze global
	fmt.Println("Esto es un print")
	var x = "hola" // declaro sin poner el tipo de dato, si doy su valor inicial

	var y string // si no doy su valor incial tengo que poner el tipo de dato
	y = "adios."

	// otra forma de declarar pero su alcance sera local (es decir alcance dentro de la funcion donde fue creada)
	z := "Esto es un string" // otra forma de declarar y se tiene que inicializar el valor, aparte no se pone el tipo de dato., se puede decir que se omite el tipo de dato

	fmt.Println(x, y, z)
}

// para ejecutar este archivo el comando seria: go run 001.go
// si queremos un archivo ejecutable para produccion debemos usar este otro comando: go build 001.go => Me crea una compilacion en binario para ser llevada a produccion
