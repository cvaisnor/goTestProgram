package main

import (
	"fmt"                       // module for printing
	"go_test_program/functions" // format: name "path" (name can be anything)
)

func main() {
	fmt.Println("This is a program that uses a few functions.")
	fmt.Println()
	// call the GreetUser function from the functions package
	name := functions.GreetUser() // global functions use PascalCase

	fmt.Println("The next line will be printed from a TestScript function.")
	functions.TestScript(name)

	fmt.Println("Would you like to convert a temperature?")
	fmt.Println("Enter 1 for Celsius to Fahrenheit, 2 for Fahrenheit to Celsius, or 3 to exit.")
	var choice int
	fmt.Scan(&choice)
	for choice != 3 {
		if choice == 1 {
			functions.GetFahrenheit() // call the GetFahrenheit function
			choice = 3
		}
		if choice == 2 {
			functions.GetCelsius() // call the GetCelsius function
			choice = 3
		}
		if choice != 1 && choice != 2 && choice != 3 {
			fmt.Println("Invalid input. Please try again.")
			fmt.Scan(&choice)
		}
	} // end for loop
	fmt.Println("Thank you for using this program.")
}
