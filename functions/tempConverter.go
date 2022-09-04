package functions // package name

import "fmt"

// global functions use PascalCase

// this is a function that converts Celsius to Fahrenheit
func GetFahrenheit() {
	fmt.Println("Enter a temperature in Celsius to convert to Fahrenheit: ")
	var celsius float64
	fmt.Scan(&celsius)
	ans := (celsius * 9 / 5) + 32
	fmt.Println("The temperature in Fahrenheit is", ans)
	fmt.Println()
}

// this is a function that converts Fahrenheit to Celsius
func GetCelsius() {
	fmt.Println("Enter a temperature in Fahrenheit to convert to Celsius: ")
	var fahrenheit float64
	fmt.Scan(&fahrenheit)
	ans := (fahrenheit - 32) * 5 / 9
	fmt.Println("The temperature in Celsius is", ans)
	fmt.Println()
}
