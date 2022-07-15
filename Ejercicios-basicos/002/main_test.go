// si el test esta en el mismo package entonces no es necesario poner en el package _test vasta con poner el mismo nombre
package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var sum = Add(3, 0)
	if sum != 3 {
		t.Error("la suma es correcta")
		t.Fail() // indico que el test fallo
	} else {
		t.Log("Test Add fincalizado correctamente") // con esto indico que el test paso la prueva
	}
}
