package main

import "fmt"

func main() {

	usuarios := make(map[string]string)

	usuarios["Karla"] = "karla@gmail.com"
	usuarios["Juan"] = ""

	// ahora para obtener el valor isamos lo siguiente
	// el segundo valor hace referencia a un boolano donde me retorna true si se puedo obtener un valor con respecto a una llave
	correo, ok := usuarios["Karla"]
	// correo2, ok := usuarios["juan"]

	if ok {
		fmt.Println(correo)
	} else {
		fmt.Println("No fue posible obtener el valor")
	}

	// otra forma de hacer el condicional
	/*
		// estructura mas adecuada siempre que vayamos obtener un valor de un mapa
		if correo, ok := usuarios["juan"]; ok {
			fmt.Println(correo)
		} else {
			fmt.Println("No fue posible obtener el valor")
		}

	*/
}
