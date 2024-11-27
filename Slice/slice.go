package main

import "fmt"

func main() {
	number := []int{1, 2, 3, 4, 5}
	number = append(number, 10, 20, 30)
	fmt.Println("Number :", number)
	fmt.Println("Value at second index is :", number[1])
	fmt.Printf("Data type: %T \n", number)

	fmt.Println("Length of array :", len(number))
	fmt.Println("Capacity of array :", cap(number))

	//make function - allow you to create a slice
	name := make([]string, 5)
	fmt.Println("Name of person are as follows:", name)
	fmt.Println("Length of array :", len(name))
	fmt.Println("Capacity of array :", cap(name))

	//in case of slice the capacity is not fixed it double the value acc to the need

}
