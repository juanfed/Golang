// otra forma de crear slices, por medio de la funcion make

// nota, los arrays son la base de cualquier slice
package main

import "fmt"

func main() {
	// recibe como parametros el tipo de dato, longitud y capacidad maxima
	slice1 := make([]int, 0, 3)

	fmt.Println("Imprecion del slice: ", slice1)
	fmt.Println("Impreción de la longitud: ", len(slice1))
	fmt.Println("Impreción de la capadidad: ", cap(slice1))

	// nota si cambio la longitud inicial de 0 a otro numero, me mostrará el array con datos, en este caso "ceros" ya que no he espesificado datos

	slice2 := make([]int, 3, 5)

	fmt.Println("Impresion del segundo slice")
	fmt.Println("Imprecion del slice: ", slice2)
	fmt.Println("Impreción de la longitud: ", len(slice2))
	fmt.Println("Impreción de la capadidad: ", cap(slice2))

	// podemos modificar los valores de los indices tal cual como lo hacemos comunmente

	slice2[0] = 100
	slice2[1] = 200
	slice2[2] = 300
	// si agrego otro elemento de esta forma me da error ya que superaria la capacidad maxima, seria ya haciendo uso del "appen"

	fmt.Println("Slice con cambios en los indices", slice2)

	// tambien podemos agregar elementos

	slice2 = append(slice2, 24)

	fmt.Println(slice2)
}
