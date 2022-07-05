// acá se veran los slices
package main

import "fmt"

func main() {
	// A diferencia de los arrays, los slices son dinamicos en cuenato al tamaño, por ende la longitud puede cambiar

	// un array no puee existir sin un arreglo

	// Se declara igual que un array solo que no ponemos nada dentro de los corchetes

	numeros := []int{1, 2, 3, 4, 5}
	fmt.Printf("Imprecion del slice: %v \n", numeros)

	// con esto podemos añadir elementos al slice
	// no modifica la variable original, solo retorna un nuevo objeto que almacenasmos en la misma variable
	numeros = append(numeros, 12)
	numeros = append(numeros, 72)

	fmt.Printf("Imprecion del nuevo slice: %v \n", numeros)

	nuevoSlice := numeros[0:5]
	tercerSlice := numeros[1:7]
	//los slices no son mas que la refencia a una porcio de un arreglo, por ende si modifico numeros asi sea despues de declarar nuevoSlice, este tambien se verá afectado

	numeros[2] = 24         //modifique el slice original
	fmt.Println(numeros)    // muestro que modifique el array original
	fmt.Println(nuevoSlice) // por ende este slice tamien se modificara independiente de donde haya echo la modificacion del slice original.

	fmt.Println("Tercer slice :", tercerSlice)

	// tenemos que tener presente lo siguiente para trabajar con los slices
	// 1) Un puntero
	// 2) Una longitud
	// 3) Una capacidad

	// para la longitud hacemos uso de "len"
	// para la capaidad se hace suo de "cap"

	meses := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "septiembre"}

	longitud := len(meses)
	capacidad := cap(meses)

	fmt.Println("La longitud del slices es: ", longitud, "\nLa pacidad del slice es: ", capacidad)

	meses = append(meses, "Octubre")
	longitud = len(meses)
	capacidad = cap(meses)

	// estos valores puden cambiar si se agregan o se eliminan datos
	fmt.Println("La longitud del slices es: ", longitud, "\nLa pacidad del slice es: ", capacidad)
	// la capacidad se duplica, esto es porque la funcion append genera un nuevo arreglo si dicho arreglo o slice se encuentra en su capadidad maxima. Para no generar objetos de forma innecesaria.

	meses = append(meses, "Noviembre", "Diciembre")

	// como no rebaso su capacidad maxima,esta seguirá siendo la misma
	fmt.Println("La capacidad es: ", cap(meses), "\nLa longitud es de: ", len(meses))

	// podemos verificar esto imprimiendo sus direcciones de memoria
}
