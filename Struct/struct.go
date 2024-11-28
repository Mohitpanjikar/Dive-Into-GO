package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	age       int
}

type phone_number struct {
	number     int
	phone_type string
}

type address struct {
	city  string
	state string
}

type employee struct {
	person_Details person
	contact_no     phone_number
	home           address
}

func main() {
	// var struct_name_var struct_name
	var person2 person
	fmt.Println(person2)

	var person1 person
	person1.firstname = "Mohit"
	person1.lastname = "Panjikar"
	person1.age = 25

	fmt.Println(person1)

	// shorthand notation for adding value into the struct -
	person3 := person{
		firstname: "Josh",
		lastname:  "Smith",
		age:       30,
	}
	fmt.Println(person3)

	// using new keyword -
	var person4 = new(person)
	person4.firstname = "Robert"
	person4.lastname = "Downey"
	person4.age = 40
	fmt.Println(*person4) // Dereference the pointer

	// Initialize employee struct with correct syntax
	employee1 := employee{
		person_Details: person{
			firstname: "Riya",
			lastname:  "Rohan",
			age:       28,
		},
		contact_no: phone_number{
			number:     1234567890,
			phone_type: "home",
		},
		home: address{
			city:  "New York",
			state: "NY",
		},
	}
	fmt.Println(employee1.person_Details)
	fmt.Println(employee1.contact_no)
	fmt.Println(employee1.home)
}
