// manejo de errores
package main

// como en go no existe las execciones, por ende tenemos que validar cualquier error
import (
	"errors" // paquete que nos ayudara en ese proceso de crear errores.
	"fmt"
)

func divicion(dividendo, divisor int) (int, error) {
	// validamos sobre el paramtro divisor
	if divisor == 0 {
		return 0, errors.New("No es posible dividir sobre 0.") // creamos un error que se le mostrará al usuario para que este sepa que esta haciendo mal
	} else {
		return dividendo / divisor, nil // usamos nil para representar la ausencia de un valor, ya que no vamos a retornar un error, se hace porque en la funcion espesificamos que retornaremos 2 cosas y no solo una. nil es igual a null de javascript
	}
}

func main() {
	if result, err := divicion(10, 0); err != nil {
		fmt.Println(err)
		// si despues del error queremos que el programa finalize, usamos la funcion panic, pasando como argumento nuestro error
		panic(err)
	} else {
		fmt.Println("El resultado es: ", result)
	}
}
