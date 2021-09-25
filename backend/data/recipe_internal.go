package data

import (
	"database/sql"
	"encoding/json"
	"foodbook/backend/data/models"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type RecipeInternal struct {
	Name         string   `json:"name"`
	Instructions []string `json:"instructions"`
}

var RECIPE_PATH = "configs/recipes/"

func readRecipes(recipeFileNames []string) ([]RecipeInternal, error) {
	var recipes []RecipeInternal
	for _, recipeFileName := range recipeFileNames {
		recipeStr, err := ioutil.ReadFile(RECIPE_PATH + recipeFileName)
		if err != nil {
			return recipes, err
		}
		recipe := RecipeInternal{}
		json.Unmarshal([]byte(recipeStr), &recipe)
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}

func InsertNewRecipeInstructions(db *sql.DB, recipe RecipeInternal) error {
	err := models.InsertRecipe(db, models.Recipe{Name: recipe.Name, Rating: 0})
	if err != nil {
		return err
	}
	stored_recipe, err := models.GetRecipe(db, recipe.Name)
	if err != nil {
		return err
	}
	for i, instruction := range recipe.Instructions {
		err = models.InsertInstruction(db, models.Instruction{RecipeId: stored_recipe.Id, Number: i, Text: instruction})
		if err != nil {
			return err
		}
	}
	return nil
}

func MaybeUpdateExistingRecipeInstructions(db *sql.DB, existing_recipe models.Recipe, new_recipe RecipeInternal) error {
	existing_instructions, err := models.GetInstructionsFromRecipe(db, existing_recipe.Id)
	if err != nil {
		log.Printf("Got error: %v when getting instructions from recipe: %v\n", err, existing_recipe.Name)
		return err
	}
	existing_instructions_by_order := make(map[int]models.Instruction)
	for _, existing_instruction := range existing_instructions {
		existing_instructions_by_order[existing_instruction.Number] = existing_instruction
	}
	for i, instruction := range new_recipe.Instructions {
		if existing_instruction, ok := existing_instructions_by_order[i]; ok && instruction != existing_instruction.Text {
			err := models.UpdateInstruction(db, models.Instruction{Id: existing_instructions_by_order[i].Id, RecipeId: existing_recipe.Id, Number: i, Text: instruction})
			if err != nil {
				log.Printf("Got error: %v when updating instruction: %v in recipe: %v\n", err, i, existing_recipe.Name)
				return err
			}
		}
	}
	return nil
}

func InsertOrUpdateRecipes(recipes []RecipeInternal) error {
	db := GetDatabaseSingleton().Db
	for _, recipe := range recipes {
		existing_recipe, err := models.GetRecipe(db, recipe.Name)
		if err != nil {
			if err == sql.ErrNoRows {
				err = InsertNewRecipeInstructions(db, recipe)
				if err != nil {
					return err
				}
			} else {
				log.Printf("Got error: %v when getting recipe: %v\n", err, recipe.Name)
				return err
			}
		} else {
			err = MaybeUpdateExistingRecipeInstructions(db, existing_recipe, recipe)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func InitializeRecipes() error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	log.Println("Recipe path: ", filepath.Join(pwd, RECIPE_PATH))
	recipeFiles, err := ioutil.ReadDir(filepath.Join(pwd, RECIPE_PATH))
	if err != nil {
		return err
	}
	recipeFileNames := []string{}
	log.Println("recipe files length: ", len(recipeFiles))
	for _, recipeFile := range recipeFiles {
		recipeFileNames = append(recipeFileNames, recipeFile.Name())
		log.Println("Recipe file name: ", recipeFile.Name())
	}
	recipes, err := readRecipes(recipeFileNames)
	if err != nil {
		return err
	}
	return InsertOrUpdateRecipes(recipes)
}
