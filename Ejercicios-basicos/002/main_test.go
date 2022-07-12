package main_test

import (
	"testing"
	main "v1/Ejercicios-basicos/002"
)

func TestAdd(t *testing.T) {
	var sum = main.Add(3, 0)
	if sum != 3 {
		t.Error("la suma es correcta")
		t.Fail() // indico que el test fallo
	} else {
		t.Log("Test Add fincalizado correctamente") // con esto indico que el test paso la prueva
	}
}
