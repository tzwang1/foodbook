package models

import "database/sql"

type Instruction struct {
	Id       string
	RecipeId string
	Text     string
}

const INSTRUCTION_TABLE_NAME = "instructions"

const INITIALIZE_INSTRUCTION_TABLE_QUERY = `
	CREATE TABLE IF NOT EXISTS ` + INSTRUCTION_TABLE_NAME + ` (
	id serial PRIMARY KEY,
	recipeId text NOT NULL,
	text integer
	CONSTRAINT fk_recipe
      FOREIGN KEY(recipeId) 
	  REFERENCES recipes(id)
	);`

func InsertInstruction(db *sql.DB, instruction Instruction) (err error) {
	sqlStatement := `INSERT INTO instructions(recipeId, text) VALUES ($1, $2);`
	_, err = db.Exec(sqlStatement, instruction.RecipeId, instruction.Text)
	return err
}
