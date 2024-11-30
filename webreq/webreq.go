package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning web service")

	// Sending an HTTP GET request to the specified URL
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		// Handling any error that occurs during the GET request
		fmt.Println("Error occurred while making GET request:", err)
		return
	}
	// Ensure the response body is closed after reading to prevent resource leaks
	defer res.Body.Close()

	// Reading the response body
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// Handling any error that occurs while reading the response
		fmt.Println("Error occurred while reading response body:", err)
		return
	}

	// Converting the byte slice to a string and printing the response
	fmt.Println("Response:", string(data))

	// Explanation:
	// 1. We first send a GET request to the API endpoint and capture the response in 'res'.
	// 2. Then, we read the response body using ioutil.ReadAll, which returns the body as a byte slice.
	// 3. Finally, we convert the byte slice to a string for easy reading and print it.
}
