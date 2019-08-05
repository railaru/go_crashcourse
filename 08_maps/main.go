package main

import "fmt"

func main() {
	//define map
	// emails := make(map[string]string)

	// //assign key values
	// emails["bob"] = "bob@gmail.com"
	// emails["joe"] = "joe@gmail.com"
	// emails["mike"] = "mike@gmail.com"

	//declare map and add key values

	emails := map[string]string{"bob": "bob@gmail.com", "joe": "joe@gmail.com", "mike": "mike@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["bob"])

	//delete from map
	delete(emails, "bob")
	fmt.Println(emails)
}
