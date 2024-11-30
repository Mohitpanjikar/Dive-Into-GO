package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Learning about files")

	file, err := os.Create("example.txt")

	if err != nil {
		fmt.Println("Error creating file:", err)
	}

	content := "i am adding file data"
	io.WriteString(file, content)

	defer file.Close()
}
