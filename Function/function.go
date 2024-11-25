package main

import "fmt"

// simpleFunc prints a simple message.
func simpleFunc() {
	fmt.Println("Hello from simple function")
}

// addTwo adds two integers and returns the result.
func addTwo(a, b int) int {
	return a + b
}

// subTwo subtracts the second integer from the first and returns the result.
func subTwo(c, d int) int {
	return c - d
}

func main() {
	fmt.Println("We are learning functions in Go")

	// Call simpleFunc
	simpleFunc()

	// Call addTwo and subTwo
	ans := addTwo(10, 20)
	sub := subTwo(10, 20)

	// Print results
	fmt.Println("Addition:", ans)
	fmt.Println("Subtraction:", sub)
}
