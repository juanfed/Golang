// un ejemplo mas real sobre programar funciones usando la palabradefer
// leeremos el contenido de un archivo .txt
// tambien aré uso de la funcion recover, el cual hace que en caso de darse un panic, no finalizar el programa de forma abructa.

package main

import (
	"fmt"
	"os"
) // este paquete me servira para leer el archivo .txt

func main() {
	file, err := os.Open("example.txt") // direccion del archivo

	if err != nil {
		panic("No fue posible obtener el archivo")
	}

	// siempre que trabajemos con archivos de nuestro sistema una vez dejado de usarlos sera necesario cerrarlos para que otro programa haga uso de ellos
	defer func() {
		fmt.Println("cerramos el archivo!!")
		file.Close() // cerrará el archivo
	}()

	// declaro una nueva variable
	contenido := make([]byte, 254)

	longitud, _ := file.Read(contenido)

	contenidoArchivo := string(contenido[0:longitud])

	fmt.Println(contenidoArchivo)
}
