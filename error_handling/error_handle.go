package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("Error handling in Go")

	// Example 1: Handle both the result and error
	ans, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", ans)
	}

	// Example 2: Use underscore to ignore the error
	ans, _ = divide(10, 0) // We ignore the error here
	fmt.Println(ans)       // But this will still print 0, as we ignored the error
}
