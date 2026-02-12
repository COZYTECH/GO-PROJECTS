package main

import (
	"fmt"
)

func main() {
// Create a map to store the ages of people
//Maps are key-value stores, like dictionaries.

contacts := map[string]string{
    "Alice": "alice@example.com",
    "Bob":   "bob@example.com",
}

// Add a new contact
contacts["Sam"] = "sam@example.com"
contacts["cozy"] = "cozy@example.com"

// Loop through map
for name, email := range contacts {
    fmt.Println(name, ":", email)
}

}