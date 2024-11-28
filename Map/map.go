package main

import "fmt"

func main() {
	/* Maps in Golang
	• In Go, a map is a data structure that provides an unordered collection of key-value pairs, where each key must be unique.
	• It is similar to dictionaries or hash maps in other programming languages.
	• Maps are used to associate values with keys and allow for efficient retrieval of values based on those keys. */
	fmt.Println("Learning about maps")

	// Creating a map
	// map_name := make(map[type_of_key]type_of_value)
	studentGrades := make(map[string]int)

	// Adding values to the map
	studentGrades["Prince"] = 100
	studentGrades["Alice"] = 90
	studentGrades["Bob"] = 85 // Corrected the key from "'Bob" to "Bob"
	studentGrades["Charlie"] = 95

	// Accessing values from the map
	fmt.Println("Marks of Bob:", studentGrades["Bob"])

	// Updating a value in the map
	studentGrades["Bob"] = 100
	fmt.Println("New marks of Bob:", studentGrades["Bob"])

	// Deleting a key-value pair from the map
	delete(studentGrades, "Bob")

	// Checking if "Bob" exists in the map before accessing it
	if grade, exists := studentGrades["Bob"]; exists {
		fmt.Println("Marks of Bob:", grade)
	} else {
		fmt.Println("Marks of Bob: Key does not exist")
	}
}
