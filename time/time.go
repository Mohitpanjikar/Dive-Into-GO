package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current time
	currentTime := time.Now()

	// Print the current time and its type
	fmt.Println("Current time: ", currentTime)
	fmt.Printf("Type of currentTime: %T\n", currentTime)

	// Format the current time in a custom format
	formatted := currentTime.Format("2006/01/02, Monday")
	fmt.Println("Formatted time: ", formatted)

	// Parse a date string into a time.Time object
	layoutStr := "2006-01-02"
	dateStr := "2023-11-25"
	formattedTime, err := time.Parse(layoutStr, dateStr)
	if err != nil {
		fmt.Println("Error parsing date:", err)
	} else {
		fmt.Println("Parsed time:", formattedTime)
	}
}

// main folder - package  main - import time
// time folder - package time
