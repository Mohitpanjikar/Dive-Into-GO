package main

import "fmt"

func main() {
	// Declare an integer variable 'num' and assign it a value
	var num int
	num = 2

	// Declare a pointer 'ptr' to hold the address of 'num'
	var ptr *int
	ptr = &num

	// Print the value of 'num' and its address using the pointer 'ptr'
	fmt.Println("num is", num)
	fmt.Println("pointer contains", ptr)               // ptr contains the memory address of 'num'
	fmt.Println("dereferenced pointer contains", *ptr) // *ptr gives the value stored at the address

	// Declare another integer variable 'num2' and pointer 'ptr2'
	num2 := 10
	ptr2 := &num2

	// Print the value of 'num2' and its address using the pointer 'ptr2'
	fmt.Println("num2 is", num2)
	fmt.Println("pointer contains", ptr2)
	fmt.Println("dereferenced pointer contains", *ptr2)

	// Declare a pointer 'ptr3' without initialization, which defaults to nil
	var ptr3 *int

	// 'ptr3' is nil, so dereferencing it will cause a runtime error.
	// To prevent this, we should check if ptr3 is nil before dereferencing
	fmt.Println("pointer contains", ptr3)
	if ptr3 != nil {
		fmt.Println("dereferenced pointer contains", *ptr3)
	} else {
		fmt.Println("cannot dereference a nil pointer")
	}
}
