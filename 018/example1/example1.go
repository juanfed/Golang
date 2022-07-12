// sentencia defer, sirve para programar funciones
// sirve para programar funcionespara ue estas sean ejecutas una vez que la funcion que las haya llamado finalicen

package main

import "fmt"

func funcion1() {
	fmt.Println("Hola soy la primera funcion")
}

func funcion2() {
	fmt.Println("Hola soy la segunda funcion")

}

func main() {
	fmt.Println("Este mensaje es directamente de la funion main")
	defer funcion1() // me movera esta linea de codigo al final, es decir se ejecutara cuando en el bloque de la funcion main finalize
	funcion2()
}

// validando errores esta palabra defer nos resultará muy util
