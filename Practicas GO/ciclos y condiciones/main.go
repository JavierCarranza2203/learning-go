package main

import "fmt"

func main() {
	miMapa := map[string]string {
		"Colombia": "Bogotá",
		"Venezuela": "Caracas",
		"Brasil": "Brasilia",
		"Chile": "Santiago",
	}

	miMapa["Argentina"] = "Buenos Aires"
	delete(miMapa, "Venezuela")
	delete(miMapa, "Colombia")
	delete(miMapa, "Argentina")

	if(len(miMapa) > 3) {
		fmt.Println("El diccionario tiene mas de 3 elementos")
	} else {
		fmt.Println("El diccionario tiene menos de 3 elementos")
	}

	fmt.Println("El mapa es: ", miMapa)
}