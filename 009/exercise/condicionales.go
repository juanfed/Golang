package main

// Ejercicios condicionales
import "fmt"

func main() {
	// programa para calcular el indice de msa corporal
	name := ""
	heigth := .0 // estatura
	weight := .0 // peso
	imc := .0    // indice de masa corporal

	fmt.Print("Ingrese su nombre: ")
	fmt.Scanf("%s\n", &name)

	fmt.Print("Ingrese su estatura en metros: ")
	fmt.Scanf("%f\n", &heigth)

	fmt.Print("Ingrese su peso corporal en Kg: ")
	fmt.Scanf("%f\n", &weight)

	imc = (weight / (heigth * heigth))

	if imc <= 18.5 {
		fmt.Print("Tiene un peso muy bajo, inferior al normal")
	} else if imc > 18.5 && imc <= 24.5 {
		fmt.Print("Tiene un peso normal 😎")
	} else if imc >= 25.0 && imc <= 29.9 {
		fmt.Print("Tiene un peso superior al normal")
	} else if imc > 30.0 {
		fmt.Print("Tiene problemas con su peso, sufre de obesidad")
	}
}
