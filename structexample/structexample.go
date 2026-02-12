package main

import "fmt"

func main() {
	// Define a struct to represent a person
	type Person struct {
		Name string
		Age  int
	}

	contact := Person{Name: "Alice", Age: 30}
	fmt.Println(contact)
}