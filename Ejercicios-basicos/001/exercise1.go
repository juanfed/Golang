package main

import (
	"fmt"
)

var number int
var numbers = []int{}

func Add() {
	for i := 0; i <= 2; i++ {
		fmt.Print("Ingrese un numero: ")
		fmt.Scanf("%d\n", &number)
		numbers = append(numbers, number)
	}
}

func main() {

	Add()

	var bigNumber = numbers[0]

	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] < numbers[i+1] {
			bigNumber = numbers[i+1]
		}
	}

	fmt.Println("El arreglo es:", numbers)
	fmt.Println("El número más grande es: ", bigNumber)
}
