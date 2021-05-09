package data

import (
	"encoding/json"
	"io/ioutil"
)

type RecipeInternal struct {
	Name         string   `json:"name"`
	Instructions []string `json:"instructions"`
}

func readRecipes(recipeFileNames []string) ([]RecipeInternal, error) {
	var recipes []RecipeInternal
	for _, recipeFileName := range recipeFileNames {
		recipeStr, err := ioutil.ReadFile(recipeFileName)
		if err != nil {
			return recipes, err
		}
		recipe := RecipeInternal{}
		json.Unmarshal([]byte(recipeStr), &recipe)
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}

func InsertOrUpdateRecipes(recipes []RecipeInternal) error {
	return nil
}

func InitializeRecipes() error {
	recipeFiles := []string{"filename1", "filename2"}
	recipes, err := readRecipes(recipeFiles)
	if err != nil {
		return err
	}
	return InsertOrUpdateRecipes(recipes)
}
