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

	// otra forma de usar este condicional switch

	switch calificacion {
	case 10:
		fmt.Print("Su calificación ha sido perfecta")
	case 8, 9:
		fmt.Print("Aprovaste la materia")
	case 7, 6:
		fmt.Print("Aprovaste a materia, pero neceistas estudiar más.")
	case 0, 1, 2, 3, 4, 5:
		fmt.Print("No aprovaste la materia")
	default:
		fmt.Print("Calificación no valida")
	}

}
