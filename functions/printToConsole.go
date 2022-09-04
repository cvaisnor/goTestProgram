package functions // package name

import "fmt"

// This is a test script
func TestScript(name string) { // global functions use PascalCase
	fmt.Println("This is a test script.")
	fmt.Println("The name you entered is", name)
	fmt.Println()
}

// this function prompts the user for their name
func getName() string { // local functions use camelCase
	fmt.Println("Enter your name: ")
	var name string
	fmt.Scan(&name)
	return name
}

// this function greets the user with their name
func GreetUser() string {
	name := getName() // call the getName function (local)
	fmt.Println("Hello,", name)
	fmt.Println()
	return name
}
