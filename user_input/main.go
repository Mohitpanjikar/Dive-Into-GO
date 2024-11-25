package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Using bufio to read a full line of input
	fmt.Println("What's your name?")

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n') // Read input until the newline character

	fmt.Println("Hello, good morning:", name)
}
