package main

import "fmt"

func main() {
	//if - else
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is smaller than 5")
	}

	//if - else if - else
	z := 20
	if z > 10 {
		fmt.Println("value of z is greater than 10")
	} else if z < 10 {
		fmt.Println("value of z is smaller than 10")
	} else {
		fmt.Println("value of z is 10")
	}
}
