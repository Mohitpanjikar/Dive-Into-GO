package main

import "fmt"

func main() {
	age := 25
	name := "Alice"
	height := 5.8234567

	// Print age, height, and name
	fmt.Println("Age:", age, "Height:", height, "Name:", name)

	// Print a simple message
	fmt.Println("Hello, world!")

	// Formatted output
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Height: %.2f\n", height)
	fmt.Printf("Type of height: %T\n", height)
}
