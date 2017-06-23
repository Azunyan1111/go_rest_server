package main

import (
	"net/http"
	"fmt"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	fmt.Print(" \n ---Hello World---\n")
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}
