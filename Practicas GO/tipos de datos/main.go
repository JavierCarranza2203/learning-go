package main

import (
	"fmt"
	"strings"
)

func main() {
	//Numeros
	entero := 10
	decimal := 3.14
	suma := entero + int(decimal)

	fmt.Println(suma)

	//Texto
	mensaje := "Hola, "
	concatenando := mensaje + "Javier"
	enMayuscula := 	strings.ToUpper(concatenando)

	fmt.Println(enMayuscula)

	//Booleanos
	esVerdadero := true

	fmt.Println(esVerdadero)

	//Arrays y Slice
	arrayFijo := [3]int{1, 2, 3}
	sliceVariable := []int{4, 5, 6}

	sliceVariable = append(sliceVariable, 7)

	fmt.Println(arrayFijo)
	fmt.Println(sliceVariable)

	//Mapas o diccionarios
	diccionario := map[string] int {
		"clave1" : 1,
		"clave2" : 2,
	}

	fmt.Println(diccionario)

	//Structs

	//NOTA: Si quiero que algo tenga modificador publico al exportar, debe tener la inicial en mayusculas
	//Similar al export / import de JavaScript
	type Persona struct {
		Nombre string //Publico (para cualquier archivo)
		Edad int //Publico (para cualquier archivo)
		contrasenia string //Privado (solo de este archivo)
	}

	persona := Persona{Nombre: "Javier", Edad: 22, contrasenia: "1234"}

	fmt.Println(persona)
}