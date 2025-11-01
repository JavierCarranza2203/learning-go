package main

import "fmt"

func main() {
	//Estructuras y modelado de datos

	//Condicional
	edad := 20
	//Usando Assertive / negative programming

	if edad < 18 {
		fmt.Println("Eres menor de edad")
		return
	}

	fmt.Println("Eres menor de edad")

	//Bucle for
	for i := 0; i < 5; i++ {
		//Print format %d = digit
		fmt.Printf("Iteración: %d\n", i)
	}

	//Bucle while
	n := 0
	for n < 3 {
		fmt.Printf("Iteración while: %d", n)
		n++
	}
}