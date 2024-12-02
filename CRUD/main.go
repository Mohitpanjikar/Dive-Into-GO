package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	// CRUD - create, read, update, delete
	fmt.Println("Learning about CRUD in Golang")

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

	//way 01 -
	/*
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error occurred while reading response body:", err)
			return
		}
		fmt.Println("Response:", string(data))
	*/

	//way 2 - Define a Todo struct to hold the decoded JSON response -
	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error occurred while decoding response body:", err)
		return
	}

	// Print the decoded Todo
	fmt.Printf("Todo: %+v\n", todo)
}
