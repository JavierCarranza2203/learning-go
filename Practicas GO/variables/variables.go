package main

import "fmt"

func main() {
	//Forma extendida de escribir una variable
	var nombrePersona string = "Javier"

	//Sin especificar tipo de dato (lo asigna dinámicamente)
	var segundoNombrePersona = "Armando"

	//Con método de asignación
	apellidoPersona := "Carranza"

	fmt.Println("Hola mundo desde Go, mi nombre es " + nombrePersona + " " + segundoNombrePersona + " " + apellidoPersona)

	/* VARIABLES NUMERICAS */
	var anioActual int16 = 2024
	var anioReducido int16 = 24
	edad := 22

	//casteo se usa int(variable)

	fmt.Println("El año actual es: ", anioActual)
	fmt.Println("El año actual es: ", anioReducido)
	fmt.Println("La edad es: ", edad)

	//No se aceptan cambios de tipos de datos como es JavaScript.

	/* ARREGLOS */
	var listaFrutas = [4] string { "Pera", "Naranja" }
	fmt.Println(listaFrutas[0])

	//Ejemplo arreglos
	// listaPaises := [3] string {"Peru", "Chile", "Brasil"}

	//Ejemplo slide / listas
	listaPaises := [] string {"Peru", "Chile", "Brasil"}

	//El metodo append agrega un elemento a una lista y regresa una nueva lista
	listaPaises = append(listaPaises, "Chile")

	//Asigna los valores desde el indice 1 hasta y para de agregar cuando el indice es 3 (no lo agrega, solo agrega 1 y 2)
	//Si no se especifica alguno, agarra desde el primero o hasta el ultimo
	listaPaises2 := listaPaises[1:3]
	
	fmt.Println(listaPaises2)
}