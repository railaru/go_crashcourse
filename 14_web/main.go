package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	template := "<div><h1>Article Title</h1><p>lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum </p></div>"
	fmt.Fprintf(w, template)
}

func about(w http.ResponseWriter, r *http.Request) {
	template := "<div><h1>About page</h1><p>lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum lorem ipsum </p></div>"
	fmt.Fprintf(w, template)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server Starting...")
	http.ListenAndServe(":3000", nil)
}
