package main

import (
	"fmt"
	"mylearning/myutils"
)

func main() {
	// Simple hello world message
	fmt.Println("Learning GoLang - Hello World")
	fmt.Println("---------------------------")
	myutils.PrintMessage("Hello world")
	fmt.Println("---------------------------")

	// Variable declarations with different types
	var name string = "Mohit Panjikar"
	var version = "1.1"
	var isPublic bool = true
	var size int = 100

	fmt.Println(name)
	fmt.Println("Version:", version)
	fmt.Println("Is Public:", isPublic)
	fmt.Println("Size:", size)

	// Integer and float declarations
	var money int = 10001
	var savingBalance = 20000
	fmt.Println("Money:", money)
	fmt.Println("Saving Balance:", savingBalance)

	var interestRate float64 = 2.2
	fmt.Printf("Interest Rate: %.2f\n", interestRate)

	// Constant declaration
	const pi = 3.1416
	fmt.Println("Pi:", pi)

	// Shorthand declaration
	personName := "Ram"
	fmt.Println("Person Name:", personName)

	// Public and private data demo
	var Public = "Data is public"
	var private = "Data is private"

	fmt.Println(private)
	fmt.Println(Public)

	/* Function visibility explanation:
	Public functions and variables (starting with a capital letter) can be accessed outside the package.
	Private functions and variables (starting with a lowercase letter) are accessible only within the package.

	Example:
	func PublicFunction() {
		fmt.Println("This is a public function.")
	}

	func privateFunction() {
		fmt.Println("This is a private function.")
	}
	*/
}
