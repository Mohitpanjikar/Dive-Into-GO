package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning about URL")

	// URL string
	myURL := "https://example.com:8080/path/to/resource?key1=value1&key2=value2"
	fmt.Printf("Type of URL: %T\n", myURL) // Print the type of myURL (string)

	// Converting the string into a URL object using net/url package
	parsedURL, err := url.Parse(myURL) // Convert string to URL object
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	fmt.Printf("Type of parsedURL: %T\n", parsedURL) // Print the type of parsedURL (URL object)

	// Accessing components of the URL
	fmt.Println("Scheme:", parsedURL.Scheme)  // Prints the scheme (e.g., https)
	fmt.Println("Host:", parsedURL.Host)      // Prints the host (e.g., example.com:8080)
	fmt.Println("Path:", parsedURL.Path)      // Prints the path (e.g., /path/to/resource)
	fmt.Println("Query:", parsedURL.RawQuery) // Prints the raw query string (e.g., key1=value1&key2=value2)

	// Modifying URL components
	parsedURL.Path = "/newPath"               // Modify the path component
	parsedURL.RawQuery = "username=iamprince" // Modify the query string

	// Constructing a URL string from the modified URL object
	newURL := parsedURL.String()    // Convert the modified URL object back to a string
	fmt.Println("New URL:", newURL) // Prints the new URL
}
