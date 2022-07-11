// variadic function (funciones que pueden permitir n cantidad de argumentos)
package main

import "fmt"

// con los 3 puntos le decimos que la funcion va recibir n cantidad de argumentos dentro
// esto me creará un arreglo donde almacenará esos argumentos

// variadic funcion
func promedioNota(calificaciones ...int) int {
	var promedio, suma int
	for _, calificaciones := range calificaciones {
		suma = suma + calificaciones
	}
	promedio = suma / len(calificaciones)
	fmt.Println("El promedio del estudiante fue de:", promedio)
	return promedio
}

func main() {
	// fmt es un tipo de Variadic function porque permite n cantidad de argumentos que deseemos
	fmt.Println("permite", "n", "cantidad de", "argumentos")

	promedioNota(8, 5, 7, 8, 3, 9, 7, 10)

	// tambien podria haber echo esto no mas:
	promedioEstudiante := promedioNota(3, 7, 8, 9, 5, 4, 3, 10)
	fmt.Println("Este otro estudiante tubo un promedio de:", promedioEstudiante)
}
