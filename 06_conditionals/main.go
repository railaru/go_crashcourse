package main

import "fmt"

func main() {

	//if/else if
	x := 10
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	}

	//else if
	color := "green"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is NOT blue or red")
	}

	//switch

	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	case "green":
		fmt.Println("color is green")
	default:
		fmt.Println("color is not: green, blue, red")
	}
}
