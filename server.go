package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Starting the server...")
	http.ListenAndServe(":3000", nil)
}
