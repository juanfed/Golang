// uso y repaso de "Switch"
package main

import "fmt"

func main() {
	var calificacion int

	fmt.Print("Ingrese una calificación: ")
	fmt.Scanf("%d\n", &calificacion)

	switch {
	case calificacion == 10:
		fmt.Print("Su calificacion ha sido perfecta")
	case calificacion < 10 && calificacion >= 8:
		fmt.Print("Aprovaste la materia")
	case calificacion < 8 && calificacion >= 6:
		fmt.Print("Aprovaste a materia, pero neceistas estudiar más.")
	case calificacion >= 0 && calificacion <= 5:
		fmt.Print("No aprovaste la materia")
	// podemos usar un default en caso de que ningulo de los casos anteriores se cumplen
	default:
		fmt.Print("Calificacion no valida")
	}

}
