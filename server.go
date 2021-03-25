package main

import (
	"example_app/backend/data"
	"fmt"
	"log"
	"net/http"
)

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to example server!")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func main() {
	data.Connect()
	http.HandleFunc("/", exampleHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
