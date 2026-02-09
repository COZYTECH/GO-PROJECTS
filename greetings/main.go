package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var name string = "Cozy"
	fmt.Print("Enter your name: ")
	//reads user input and stores it in the variable name
	fmt.Scanln(&name)
    greetings := []string{
        "Hi, %v! Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Pick a random greeting
    msg := greetings[rand.Intn(len(greetings))]

    // Print it
    fmt.Printf(msg+"\n", name)
    //fmt.Println("Hello, " + name + "!")
}
