package models

import "database/sql"

type Instruction struct {
	Id       string
	RecipeId string
	Number   int
	Text     string
}

const INSTRUCTION_TABLE_NAME = "instructions"

const INITIALIZE_INSTRUCTION_TABLE_QUERY = `
	CREATE TABLE IF NOT EXISTS ` + INSTRUCTION_TABLE_NAME + ` (
	id serial PRIMARY KEY,
	recipeId integer NOT NULL,
	number integer,
	text integer,
	CONSTRAINT fk_recipe
      FOREIGN KEY(recipeId) 
	  REFERENCES recipes(id)
	);`

func InsertInstruction(db *sql.DB, instruction Instruction) (err error) {
	sqlStatement := `INSERT INTO instructions(recipeId, number, text) VALUES ($1, $2);`
	_, err = db.Exec(sqlStatement, instruction.RecipeId, instruction.Text)
	return err
}

func UpdateInstruction(db *sql.DB, instruction Instruction) (err error) {
	sqlStatement := `
	UPDATE instructions SET
	text = $1
	number = $2
	WHERE id = $3;`
	_, err = db.Exec(sqlStatement, instruction.Text, instruction.Number, instruction.Id)
	return err
}

func DeleteInstruction(db *sql.DB, id string) (err error) {
	sqlStatement := `
	DELETE FROM instructions WHERE
	id = $1;`
	_, err = db.Exec(sqlStatement, id)
	return err
}

func GetInstruction(db *sql.DB, id string) (Instruction, error) {
	sqlStatement := `
	SELECT * FROM instructions WHERE id = $1;`
	row := db.QueryRow(sqlStatement, id)
	var instruction Instruction
	switch err := row.Scan(&instruction.Id, &instruction.RecipeId, &instruction.Text, &instruction.Number); err {
	case nil:
		return instruction, nil
	default:
		return Instruction{}, err
	}
}

func GetInstructionsFromRecipe(db *sql.DB, recipe_id string) ([]Instruction, error) {
	sqlStatement := `
	SELECT * FROM instructions WHERE id = $1;`
	var instructions []Instruction
	rows, err := db.Query(sqlStatement, recipe_id)
	if err != nil {
		return instructions, err
	}

	for rows.Next() {
		var instruction Instruction
		if err := rows.Scan(&instruction.Id, &instruction.RecipeId, &instruction.Text, &instruction.Number); err != nil {
			return instructions, err
		}
		instructions = append(instructions, instruction)
	}
	return instructions, nil
}
