package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Initialize an integer variable
	var num int = 42
	fmt.Println("Number is:", num)      // Print the integer value
	fmt.Printf("Data type: %T \n", num) // Print the data type of the variable

	// Converting the int data to float data
	var data float64 = float64(num)      // Type casting int to float64
	fmt.Println("Number is:", data)      // Print the float value
	fmt.Printf("Data type: %T \n", data) // Print the data type of the variable (float64)

	// Converting the int data to string data using strconv.Itoa function
	value := 123                        // Initialize another integer value
	str := strconv.Itoa(value)          // Convert int to string
	fmt.Println("Number is:", str)      // Print the string value
	fmt.Printf("Data type: %T \n", str) // Print the data type of the variable (string)

	// Converting the string data to int data using strconv.Atoi function
	str = "123"                           // Initialize a string value
	value, _ = strconv.Atoi(str)          // Convert string to int (ignoring the error)
	fmt.Println("Number is:", value)      // Print the integer value
	fmt.Printf("Data type: %T \n", value) // Print the data type of the variable (int)
}
