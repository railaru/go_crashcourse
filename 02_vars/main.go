package main

import "fmt"

func main() {

	//using var

	//var name = "Rutenis"
	var age int32 = 21
	var isCool = true
	isCool = false
	//var size float32 = 2.3
	//const size float32 = 2.3

	//shorthand
	size := 1.3
	// name := "Rutenis"
	// email := "rutenis@gmail.com"
	name, email := "Rutenis", "rutenis@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", isCool) // check type
	fmt.Printf("%T\n", name)   // check type
	fmt.Printf("%T\n", size)   // check type
}
