package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

func main() {

	var name string = "Cozy"
	fmt.Print("Enter your name: ")
	//reads user input and stores it in the variable name
	fmt.Scanln(&name)

    // Pick a random greeting
    // msg := greetings[rand.Intn(len(greetings))]

    // // Print it
    // fmt.Printf(msg+"\n", name)
    // //fmt.Println("Hello, " + name + "!")

	message, err := Hello(name)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(message)
}
func Hello(name string) (string, error) {
    if name == "" {
        return "", errors.New("empty name provided")
    }

    greeting := randomGreeting()
    return fmt.Sprintf(greeting, name), nil
}

func randomGreeting() string {
    greetings := []string{
        "Hi, %v! Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
        "Howdy, %v!",
        "Hello there, %v!",
    }

    return greetings[rand.Intn(len(greetings))]
}
