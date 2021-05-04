package main

import (
	"example_app/backend/data"
	"example_app/backend/data/models"
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

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Adding a user...")
	database := data.GetDatabaseSingleton()
	testUser := models.User{Name: "test_name", Age: 1, Email: "test@mail.com"}
	models.InsertUser(database.Db, testUser)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Getting a user...")
	database := data.GetDatabaseSingleton()
	storedUser, err := models.GetUser(database.Db, "test@mail.com")
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, "Got a user with id: %v, name: %v, age: %v, and email: %v",
		storedUser.Id, storedUser.Name, storedUser.Age, storedUser.Email)
}

func main() {
	http.HandleFunc("/", exampleHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/addUser", addUserHandler)
	http.HandleFunc("/getUser", getUserHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
