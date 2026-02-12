package main

import (
	"fmt"
)

func main(){
	// // Declare an array of integers with a length of 5
	// var numbers [5]int       // array of 5 integers
	// numbers[0] = 10
	// numbers[1] = 20
	// fmt.Println(numbers)     // [10 20 0 0 0]


	//slice
		// fruits := []string{"apple", "banana", "orange"}
		// fruits = append(fruits, "mango") // add item
		// fmt.Println(fruits)               // [apple banana orange mango]

		// for i, f := range fruits {
		// 	fmt.Println(i, f)
		// }

		fruits := []string{"apple", "banana", "orange"}
		fruits =  append(fruits, "mango")
		fmt.Println(fruits)

		for i, f := range fruits {
			fmt.Println(i, f)
		}

}