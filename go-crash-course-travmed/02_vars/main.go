package main

import "fmt"

func main() {
	// MAIN TYPES
	// string*
	// bool*
	// int*
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr - no negative no.s
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64*
	// complex64 complex128 - very large no.s

	// Using var
	// var name string = "Bob"
	name := "Bob"
	var age int = 37
	const isCool = true
	var size float32 = 1.3
	fullname, email := "Q", "q@q.com"

	fmt.Println(name, "is", age)
	fmt.Printf("%s is a type of %T\n", name, name)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", size)
	fmt.Println(fullname, age, isCool, size, email)
}
