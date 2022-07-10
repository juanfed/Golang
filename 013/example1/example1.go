// funciones en golang
package main

import "fmt"

// crearemos una funcion nueva, para ello hacemo uso de la palabra reservada func

// dentro de los parestesis van los paremos que reciebe esa funcion aparte de decir que tipo de valor sera el dato que recibirá
func saludar(username string) {
	// aca dentro irá todo el codigo que la funcion vaya a ejecutar.
	fmt.Println("Hola", username, "mucho gusto")
}

func main() {
	//aca dentro es donde llameramos a la funciones que nosotros creemos
	// en caso de que las función necesite paremetros, estos van metidos dentro de los parentesis
	saludar("Fernando")
}
