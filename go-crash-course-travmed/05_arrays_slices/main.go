package main

import "fmt"

func main() {
	// Arrays
	// var fruitArr [2]string

	// Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and assign - Shorter version of the above // Arrays & // Assign values
	// fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[1])

	fruitSlice := []string{"Apple", "Banana", "Grapes", "Kiwi"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[2:4])
}
