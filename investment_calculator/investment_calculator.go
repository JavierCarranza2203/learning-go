package main

import (
	"fmt"
	"math"
)

// ======================== CONVERSION DE TIPOS ===================================
// var futureValue = float64(investmentAmount) * math.Pow((1 + expectedReturnRate / 100), float64(years))

//Declaracion de variables en la misma linea. Si se quita el tipo de dato explicito, se pueden declarar de diferentes tipos
// Si el tipo de dato esta explicito, todas las variables son de ese tipo
//Se usa cuando el tipo de dato inferido es correcto (:=)

const inflationRate = 2.5

func main() {
	var investmentAmount, years, expectedReturnRate float64

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	//Para asignar el valor de la consola al valor de la variable se usa el & para pasar un puntero (por referencia)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)

	// fmt.Println("Future Value:", futureValue)
	//Formatting text %.(numerOfDecimals)f
	// fmt.Printf("Future Value: %.2f\n", futureValue)
	// fmt.Printf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)
	
	//fmt.Printf(`Future Value: %.2f
	//Future Value (adjusted for inflation): %.2f`, futureValue, futureRealValue) //Multiple lines string


	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow((1 + expectedReturnRate / 100), years)
	rfv = fv / math.Pow( 1 + inflationRate / 100, years)

	return fv, rfv
}