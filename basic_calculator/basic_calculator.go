package main

import "fmt"

func main() {
	var option int32
	var number1, number2 int32

	fmt.Println("Basic calculator.")
	fmt.Println("Choose an option:")
	fmt.Println("1-Add\n2-Substract\n3-Divide\n4-Multiply")
	fmt.Scan(&option)

	fmt.Print("Enter the first number: ")
	fmt.Scan(&number1)
	fmt.Println()
	fmt.Print("Enter the second number: ")
	fmt.Scan(&number2)
	fmt.Println()

	switch(option){
		case 1:
			fmt.Printf("Add = %v", number1 + number2)
		case 2:
			fmt.Printf("Substraction = %v", number1 - number2)
		case 3:
			if number2 == 0 {
				fmt.Println(fmt.Errorf("Cannot divide by zero"))
			} else {
				fmt.Printf("Divide: %.2f", float32(number1) / float32(number2))
			}
		case 4:
			fmt.Printf("Mutiply: %v", number1 * number2)
		default:
			fmt.Println(fmt.Errorf("Option is no defined"))
	}
}