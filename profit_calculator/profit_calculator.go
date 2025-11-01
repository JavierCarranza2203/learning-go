package main

import (
	"errors"
	"fmt"
	"os"
)

//Go copia los argumentos, cuando se pasa un parametro no se pasa por referencia, se pasa una copia.
//Para modificar el valor de una variable dentro de una funcion hay que pasar la direccion de memoria &
//Y en la funcion definir que vamos a recibir un puntero *

func main () {
	revenue, err := getUserInput("Revenue: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Expenses: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Tax rate: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateValues(revenue, expenses, taxRate)

	if err := writeValuesToFile(ebt, profit, ratio); err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Printf("%.2f\n", ebt)
	fmt.Printf("%.2f\n", profit)
	fmt.Printf("%.2f\n", ratio)
}

func getUserInput(text string) (float64, error) {
	var value float64
	fmt.Print(text)
	fmt.Scan(&value)

	if value <= 0 {
		return 0, errors.New("cannot use negative values or zero")
	}

	return value, nil
}

func calculateValues(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate / 100)
	ratio = ebt / profit

	return //regresa los valores aunque no los escriba porque los definí antes
}

func writeValuesToFile(ebt, profit, ratio float64) error {
	balanceText := "EBT: " + fmt.Sprint(ebt) + " PROFIT: " + fmt.Sprint(profit) + " RATIO: " + fmt.Sprint(ratio)
	return os.WriteFile("values.txt", []byte(balanceText), 0644)
}
