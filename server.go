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

func addRecipeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Adding a recipe...\n")
	database := data.GetDatabaseSingleton()
	testRecipe := models.Recipe{Name: "test_name", Rating: 0}
	models.InsertRecipe(database.Db, testRecipe)
}

func getRecipesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Getting a recipe...\n")
	database := data.GetDatabaseSingleton()
	recipes, err := models.GetRecipes(database.Db, "test_name")
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, "Got %v recipes.", len(recipes))
	for _, recipe := range recipes {
		fmt.Fprintf(w, "Got a recipe with id: %v, name: %v, and rating: %v",
			recipe.Id, recipe.Name, recipe.Rating)
	}
}

func updateRecipeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Updating a recipe...\n")
	database := data.GetDatabaseSingleton()
	newRecipe := models.Recipe{Id: "1", Name: "updatedRecipename", Rating: 5}
	err := models.UpdateRecipe(database.Db, newRecipe)
	if err != nil {
		log.Println(err)
	}
}

func deleteRecipeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete a recipe...\n")
	database := data.GetDatabaseSingleton()
	err := models.DeleteRecipe(database.Db /*id=*/, "1")
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, "Deleted a recipe with id: %v", "1")
}

func main() {
	http.HandleFunc("/", exampleHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/addRecipe", addRecipeHandler)
	http.HandleFunc("/getRecipes", getRecipesHandler)
	http.HandleFunc("/updateRecipe", updateRecipeHandler)
	http.HandleFunc("/deleteRecipe", deleteRecipeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
