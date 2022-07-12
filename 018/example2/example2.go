// un ejemplo mas real sobre programar funciones usando la palabradefer
// leeremos el contenido de un archivo .txt
// tambien aré uso de la funcion recover, el cual hace que en caso de darse un panic, no finalizar el programa de forma abructa.

package main

import (
	"fmt"
	"os"
) // este paquete me servira para leer el archivo .txt

func main() {
	defer func() {
		// esta funcion nos retornará un posible error, evitará que se mande el error y mas bien mandara un mensaje de lo ocurrido
		if err := recover(); err != nil {
			// aca entro puedo mostrar un mensaje o jecutar todas las sentencias que quiera
			fmt.Println("Ups!. El programa no pudo finalizar de forma correcta")
		}
	}()

	if file, err := os.Open("example2.txt"); err != nil {
		panic("No fue posible obtener el archivo") // lanzara un error de status 2 y finalizara de manera abructa
	} else {

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

}
