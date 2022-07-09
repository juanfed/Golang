// trabajando con for
package main

import "fmt"

func main() {
	// implementacion del ciclo for
	// El ciclo for admite de 0 a 3 bloques
	// for <variable>; <condicion>; <incremento> {sentencias a ejecutar para cada iteracion}

	// Ejemplo:
	// siempre que conoscamos la cantidad de iteraciones, podemos hacer uso de este tipo de for
	for i := 1; i <= 100; i++ {
		//fmt.Println("El valor de i es: ", i)

		if i%2 == 0 {
			fmt.Println("El numero: ", i, "es un número par")
		}

	}

	// se puede usar el for como un while
	// si no conocemos la cantidad de iteraciones a reliazar, podemos usar esta estructura para los for
	numero := 123
	contador := 0

	for numero > 0 {
		numero = numero / 10
		contador++
	}

	fmt.Println("La cantidad de digitos es: ", contador)

	// usano el for como ciclo for each
	animales := [...]string{"perro", "gato", "pez", "conejo", "vaca"}

	// primero obtenermos los elementos mediante inices y luego mediante un ciclo for each
	for indice := 0; indice < len(animales); indice++ {
		fmt.Println(animales[indice])
	}
	// ahora obtendremos los elementos como si fuera un for each
	// si no requiero el indice2 o elemento, puedo poner un "_", para evitar que el compilador me marque un error
	for indice2, elemento := range animales {
		fmt.Println("El indice es: ", indice2, "el valor es: ", elemento)
	}

	// palabra reservada breack y continue

	for x := 1; x <= 10; x++ {
		// el numero 5 no aparecera en consola
		if x == 5 {
			// esta palabra finaliza la iteraccion actual para pasar a la 5, es decir esta ejecucion se omite
			continue
		}
		if x == 7 {
			// finaliza la ejecicion del ciclo actual, por ende las siguientes ejecuciones del ciclo ya no se ejecutaran
			break
		}
		fmt.Println(x)
	}

	//emulacion del ciclo do while por medio del for
	var number = 1
	// en caso de que el numero sea 10 y queremos que por lo menos se ejecute la sentencia una vez, tal como se hace en un do while tradicional
	for ok := true; ok; ok = number < 10 { // emulo un do while
		fmt.Println("numero: ", number)
		number++
	}

	// for como ciclo infinito "aunque no es muy recomendado usar esto", por ende lo comento, podemos panarlo con panic o breack cuando se cumpla una condicion dentro del ciclo
	/*
		for {
			esto aca sera un ciclo infinito, el cual no finalizará
		}
	*/

	// funcion "panic"

	// esta funcion es usada para finalizar el programa de manera abructa cuando exista un error.

	// ejemplo programa que permite ralizar una pequeña divicion

	var dividendo, divisor int

	fmt.Print("ingresa un valor para el dividendo: ")
	fmt.Scanf("%d\n", &dividendo)

	fmt.Print("Ingresa un valor para el divisor: ")
	fmt.Scanf("%d\n", &divisor)
	// para evitar el error por divicion por cero
	// finaliza el programa pero me dira por que fue el error.
	if divisor == 0 {
		panic("No es posible la dicion por 0")
	}

	resultado := dividendo / divisor

	fmt.Println("El resultado de la divicion es: ", resultado)
}
