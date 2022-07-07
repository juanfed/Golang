// explicacion de los "mapas"

package main

import "fmt"

func main() {
	// los map son colecciones desosrganizadas de llave valor, "arreglos clave valor"

	// en make ira el tipo de dato de la clave, y luego fuera de los corchetes el tipo de dato del valor asociado a esa llave
	dias := make(map[int]string)

	// clave -- valor
	dias[0] = "Domingo"
	dias[1] = "Lunes"
	dias[2] = "Martes"
	dias[3] = "Miercoles"
	dias[4] = "Jueves"
	dias[5] = "Viernes"
	dias[6] = "Sabado"

	fmt.Println(dias)

	// para extraer valores, lo hago mediante la llave (clave)
	fmt.Println(dias[4]) // imprimo el valor almacenado en ese indice, llave o clave

	// Tambien puedo cambiar los valores que estan asociados a una llave
	dias[4] = "JUEVES"

	// para borrar una llave con su correspondiente valor hacemos uso de la funcion delete
	delete(dias, 5)
	fmt.Println(dias)

	// tambien podemos generar estructuras mas complehas como mapas que almacenes arreglos o slices

	usuarios := make(map[string][]int)

	usuarios["Juan"] = []int{10, 9, 8, 10, 5}
	usuarios["Yessica"] = []int{5, 8, 4, 10, 7}

	// Impresion de un map que almacena clave usuarios y valor slices con sus notas
	fmt.Println(usuarios)

	// otra forma en la cual se puede crear un map seria de la siguiente forma:

	nuevoMap := map[int]string{} // sin necesidad de usar la funcion make

	nuevoMap[0] = "Usuario"
	nuevoMap[1] = "Usuario 1"
	nuevoMap[2] = "Usuario 2"
	nuevoMap[3] = "Usuario 3"
	nuevoMap[4] = "Usuario 4"

	fmt.Println(nuevoMap)

	// ahora la forma para iterar sobre los mapas seria la sieguiente
	// for variable1 (llave de los registros) variable2 (valor de esa llave)

	// si la impresion es de un orden diferente es porque las llaves no poeen un orden por ende la imprecion varia de una ejecucion u otra
	for llave, valor := range nuevoMap {
		fmt.Println(llave, valor)
	}

}
