package main

import (
	"fmt"
	"strings"
)

func main() {
	println("learning about string function")

	//split - split the string by a delimiter
	data := "apple,banana,orange"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	data2 := "Mohit.love.having.fun"
	parts2 := strings.Split(data2, ".")
	fmt.Println(parts2)

	//count - count the number of times a substring appears in a string
	str := "one two two three three three"
	count_value := strings.Count(str, "three")
	fmt.Println(count_value)

	//removing the wide space - trim space
	str_trim_space := "   hello    world   "
	trimmed := strings.TrimSpace(str_trim_space)
	println(trimmed)

	//join function - join the elements of a slice into a string
	str_name := "Mohit Panjikar"
	str_phone := "1234567890"
	result := strings.Join([]string{str_name, str_phone}, " ")
	fmt.Println(result)

	//uppercase and lowercase function
	str2 := "Hello, World!"
	uppercase := strings.ToUpper(str2)
	lowercase := strings.ToLower(str2)
	fmt.Println(uppercase)
	fmt.Println(lowercase)

}
