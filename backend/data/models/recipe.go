package models

import (
	"database/sql"
)

type Recipe struct {
	Id     string
	Name   string
	Rating int
}

const RECIPE_TABLE_NAME = "recipes"

const INITIALIZE_RECIPE_TABLE_QUERY = `
	CREATE TABLE IF NOT EXISTS ` + RECIPE_TABLE_NAME + ` (
	id serial PRIMARY KEY,
	name text NOT NULL,
	rating integer)`

func InsertRecipe(db *sql.DB, recipe Recipe) (err error) {
	sqlStatement := `INSERT INTO recipes(name, rating) VALUES ($1, $2);`
	_, err = db.Exec(sqlStatement, recipe.Name, recipe.Rating)
	return err
}

func UpdateRecipe(db *sql.DB, recipe Recipe) (err error) {
	sqlStatement := `
	UPDATE recipes SET
	name = $1,
	rating = $2
	WHERE id = $3;`
	_, err = db.Exec(sqlStatement, recipe.Name, recipe.Rating, recipe.Id)
	return err
}

func DeleteRecipe(db *sql.DB, id string) (err error) {
	sqlStatement := `
	DELETE FROM recipes WHERE
	id = $1;`
	_, err = db.Exec(sqlStatement, id)
	return err
}

func GetRecipe(db *sql.DB, name string) (Recipe, error) {
	sqlStatement := `
	SELECT * FROM recipes WHERE name = $1;`
	row := db.QueryRow(sqlStatement, name)
	var recipe Recipe
	switch err := row.Scan(&recipe.Id, &recipe.Name, &recipe.Rating); err {
	case nil:
		return recipe, nil
	default:
		return Recipe{}, err
	}
}

func GetRecipes(db *sql.DB, name string) ([]Recipe, error) {
	sqlStatement := `
	SELECt * FROM recipes where name LIKE $1`
	var recipes []Recipe
	rows, err := db.Query(sqlStatement, name)
	if err != nil {
		return recipes, err
	}

	for rows.Next() {
		var recipe Recipe
		if err := rows.Scan(&recipe.Id, &recipe.Name, &recipe.Rating); err != nil {
			return recipes, err
		}
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}
