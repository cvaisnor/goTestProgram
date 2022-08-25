package main

import "fmt"

func getFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func main() {
	fmt.Println("This is a program that converts Celsius to Fahrenheit.")
	fmt.Println("Enter a temperature in Celsius: ")
	var celsius float64
	fmt.Scan(&celsius)

	fmt.Println("The temperature in Fahrenheit is: ", getFahrenheit(celsius))
}
