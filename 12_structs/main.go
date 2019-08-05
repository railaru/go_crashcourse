package main

import (
	"fmt"
	"strconv"
)

//define person struct

type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

//greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

//hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

//getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "male" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// init person using struct

	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "Female", age: 25}
	person2 := Person{"Bob", "Johnson", "Boston", "Male", 21}

	fmt.Println(person1.greet())
	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	person1.getMarried("Williams")
	fmt.Println(person1.greet())

	person2.getMarried("Thompson")
	fmt.Println(person2.greet())

}
