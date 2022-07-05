// arreglos

// los arreglos en golang son un poco diferentes a son en otros lenguajes.

package main

import "fmt"

func main() {
	// los arrays solo almacenan el tipo de dato declarado al crearlo, no almacena varios tipos de datos
	var numeros [5]int // nombre de la variable junto con el tamaño del array, si no le pongo datos al array estos por defectos seran ceros

	array1 := [5]int{100, 200, 300, 400, 500} // forma de inicializar un array
	fmt.Println(len(numeros))                 // me muestra el tamaño del array
	fmt.Println(numeros)                      // me mostrará el array que al no definir numeros seran puros ceros
	fmt.Printf("Array con valores inicializados: %v \n", array1)

	numeros[0] = 1 // forma de agregar valores a los arrays
	numeros[1] = 2
	numeros[2] = 3
	numeros[3] = 4
	numeros[4] = 12

	fmt.Println(numeros)

	// tabien podemos omitir la longitud del arreglo. Ejemplo

	array2 := [...]int{12, 23, 45, 67, 1}
	fmt.Printf("Array sin definir la longitud %v", array2)
	// en el array que no definimos la longitud, el lenguaje como tal le da una longitud que corresponderá a la cantidad de valores almacenados
	fmt.Println("Longitud del arreglo sin logitd efinida: ", len(array2))

	// Más cosas sobre los arrays

	monedas := [...]string{"Peso Colombiano", "Dolar", "Euro", "Yuan Chino"}
	fmt.Println(monedas)
	fmt.Println(monedas[2]) // acá la imprecion es por posicion donde se encuentre ese valor en el array

	// pero tambien podemos crear un array clave valor

	monedas2 := [...]string{3: "Peso Colombiano", 1: "Dolar", 2: "Euro", 4: "Yen"}

	fmt.Println(monedas2)
	// aca se llaman por los indices asi que si cambio el numeros de los indices me dara otro ersultado asosiado a ese indice mas no a su posicion
	fmt.Println(monedas2[0])
	fmt.Println(monedas2[1])
	fmt.Println(monedas2[2])
	fmt.Println(monedas2[3]) // string vacio ya que no asignamos nada a esos indices
	fmt.Println(monedas2[4])

	// tambien podemos generar un sub arreglo.

	submonedas := monedas2[0:3] // Puedo extraer de forma de array valores del array original
	// como nota cabe destacar que estos sub arreglos son slices.  Y se necesita un indice inicial y final

	fmt.Println(submonedas)
}
