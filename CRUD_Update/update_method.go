package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func getrequest() {
	// Making a GET request to fetch a todo item
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error occurred while making GET request:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error occurred while making GET request:", res.Status)
		return
	}

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error occurred while decoding response body:", err)
		return
	}

	// Print the decoded Todo
	fmt.Printf("Todo: %+v\n", todo)
}

// sending the data into the server
func performPostRequest() {
	todo := Todo{
		UserID:    23,
		Title:     "Prince Kumar",
		Completed: true,
	}

	// Convert the Todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling : ", err)
		return
	}

	// convert json data to string
	jsonString := string(jsonData)

	// convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	myURL := "https://jsonplaceholder.typicode.com/todos"

	// send POST request
	res, err := http.Post(myURL, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending request : ", err)
		return
	}

	defer res.Body.Close()

	// data, _ := ioutil.ReadAll(res.Body)
	// fmt.Println("Response : ", string(data))

	fmt.Println("Response  status : ", res.Status)
}

func performsUpdateRequest() {
	todo := Todo{
		UserID:    23,
		Title:     "Prince Kumar",
		Completed: true,
	}

	// Convert the Todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling : ", err)
		return
	}

	// convert json data to string
	jsonString := string(jsonData)

	// convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	myURL := "https://jsonplaceholder.typicode.com/todos/1"

	//craeting a Put Request -
	req, err := http.NewRequest(http.MethodPost, myURL, jsonReader)
	if err != nil {
		fmt.Println("Error sending request : ", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	//sending the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request : ", err)
		return
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response : ", string(data))
	fmt.Println("Response  status : ", res.Status)
}
func main() {
	// CRUD - create, read, update, delete
	fmt.Println("Learning about CRUD in Golang")
	getrequest()
	performPostRequest()
	performsUpdateRequest()
}

/*
This Go function, performPostRequest(), performs a POST request to the URL https://jsonplaceholder.typicode.com/todos.
It creates a Todo struct with sample data, marshals it into JSON, and then sends it as the request body.
 The function handles errors during JSON conversion and the HTTP request, and finally prints the HTTP response status.*/
