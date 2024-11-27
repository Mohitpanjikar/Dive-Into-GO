package main

import (
	"fmt"
)

func main() {
	fmt.Println("Learning about arrays")

	var name [5]string
	// var variable_name[variable_size] type

	name[0] = "Mohit"
	name[1] = "Panjikar"
	fmt.Println("Names of the person are as follows:", name)

	//putting value when we decalare -
	var number = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers are as follows:", number)
	fmt.Println("length of array :", len(number))
	fmt.Println("value at second index is :", number[1])

	var price [5]int
	fmt.Println("price is :", price)

	var price_value [5]bool
	fmt.Println("price is :", price_value)

	var price_string [5]string
	fmt.Println("price is :", price_string)
}
