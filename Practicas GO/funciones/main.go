package main

import (
	"fmt"
	"errors"
)

// func suma(a, b int) int {
// 	return a + b 
// }


func imprimirPares(numeros... int) {
	for _, numero := range numeros {
		if numero % 2 == 0 {
			fmt.Println(numero)
		}
	}
}

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		fmt.Errorf("No se puede dividir entre cero")
		return 0, errors.New("No se puede dividir entre cero")
	}

	cociente := a / b

	return cociente, nil
}

func imprimirNombres(nombres... string) {
	for _, nombre := range nombres {
		fmt.Println(nombre)
	}
}

func main() {
	fmt.Println(dividir(10,0))

	imprimirNombres("Javier", "Daniela")
}