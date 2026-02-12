package main

import "fmt"

type Person struct {
		Name string
		Age  int
	}

func (p Person) Greetings() string {
	return fmt.Sprintf("Hello, %s! You are %d years old.", p.Name, p.Age)
}

func main() {
	contact := Person{Name: "CozyTech", Age: 24}
	fmt.Println(contact.Greetings())
}