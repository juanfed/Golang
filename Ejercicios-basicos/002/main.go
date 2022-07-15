package main

import "fmt"

var option int

var num1, num2 int

func Add(number1, number2 int) int {
	return number1 + number2
}

// func Sub(number1, number2 int) int {
// 	return number1 - number2
// }

func main() {
	fmt.Print("Selecciona 1 para realizar una suma 2 para relizar una resta: ")
	fmt.Scanf("%d\n", option)

	fmt.Print("Ingresa el primer numero: ")
	fmt.Scanf("%d\n", num1)
	fmt.Print("Ingresa el segundo numero: ")
	fmt.Scanf("%d\n", num2)

	switch option {
	case 1:
		{
			Add(num1, num2)
		}
	default:
		{
			fmt.Println("Ingreso una obcion no valida")
		}
	}
}
