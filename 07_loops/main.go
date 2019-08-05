package main

import "fmt"

func main() {

	// long for loop method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	//short for loop method
	for index := 0; index <= 10; index++ {
		fmt.Printf("Number: %d\n", index)
	}

	//fizz buzz
	for i := 1; i < 100; i++ {
		if i%15 == 0 {
			fmt.Println("fizz buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
