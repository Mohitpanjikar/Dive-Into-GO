package main

import (
	"fmt"
)

func main() {
	/*	The defer statement delays the execution of the specified function until the surrounding function
		(in this case, main) returns. So “Middle of the program” will be printed just before main finishes,
		after the “Ending of the program” statement.*/

	fmt.Println("Starting of the program")

	defer fmt.Println("Middle of the program")

	fmt.Println("Ending of the program")

	//If you have two defer statements in Go, they will execute in reverse order (last-in, first-out, or LIFO) when the function returns.
	fmt.Println("Start of the function")

	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")

	fmt.Println("End of the function")
}
