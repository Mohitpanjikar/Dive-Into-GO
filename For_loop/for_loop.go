package main

import "fmt"

func main() {
	//there is only one kind of loop in go which is for loop
	for i := 0; i < 5; i++ {
		fmt.Println("number is :", i)
	}
	// break statement -
	fmt.Println("break statement")
	for i := 0; i < 5; i++ {
		if i == 2 {
			break
		}
		fmt.Println("number is :", i)
	}
	// continue statement -
	fmt.Println("continue statement")
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("number is :", i)
	}
	// range keyword - range return two things which are as follows - index and value
	number := []int{75, 21, 73, 45, 45}
	for index, value := range number {
		fmt.Printf("index is %d and value is %d\n", index, value)
	}

	date := "hello world"
	for index, value := range date {
		fmt.Printf("index is %d and char value is %c\n", index, value)
	}
}
