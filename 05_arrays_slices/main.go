package main

import "fmt"

func main() {
	//arrays

	// var fruitArr [2]string

	// //assign values
	// fruitArr[0] = "apple"
	// fruitArr[1] = "orange"

	//declare & assign

	// fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[1])

	//slices
	fruitSlice := []string{"Apple", "Orange", "Grape", "Satsuma"}
	fmt.Println(fruitSlice)
	fmt.Println(fruitSlice[1:3])
	fmt.Println(len(fruitSlice))

}
