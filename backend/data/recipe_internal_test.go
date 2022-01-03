package data

import "testing"

func TestReadRecipes(t *testing.T) {
	recipeFileNames := []string{"testrecipe.json"}
	recipes, err := readRecipesFromConfigs(recipeFileNames)
	if err != nil {
		t.Errorf("Got an unexpected error: %v\n", err)
	}
	if len(recipes) != 1 {
		t.Errorf("Got an unexpected number of recipes: %v\n", len(recipes))
	}
	if recipes[0].Name != "testName" {
		t.Errorf("Got an unexpected recipe name: %v\n", recipes[0].Name)
	}
	if len(recipes[0].Instructions) != 2 {
		t.Errorf("Got an unexpected number of recipe instructions: %v\n", len(recipes[0].Instructions))
	}
	if recipes[0].Instructions[0] != "instruction1" {
		t.Errorf("Got an unexpected recipe instruction: %v\n", recipes[0].Instructions[0])
	}
	if recipes[0].Instructions[1] != "instruction2" {
		t.Errorf("Got an unexpected recipe instruction: %v\n", recipes[0].Instructions[1])
	}
}
