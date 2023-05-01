package main

import "fmt"

func main() {
	x := 5
	y := 10

	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// else if
	color := "red"

	if color == "red" {
		fmt.Println("The Color is Red")
	} else if color == "blue" {
		fmt.Println("The Color is Blue")
	} else {
		fmt.Println("The Color is NOT Blue or Red")
	}

	// Switch

	color2 := "yellow"

	switch color2 {
	case "red":
		fmt.Println("The Color2 is Red")
	case "blue":
		fmt.Println("The Color2 is Blue")
	default:
		fmt.Println("The Color2 is NOT Blue or Red")
	}
}
