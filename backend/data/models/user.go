package models

import (
	"database/sql"
)

type User struct {
	id    string
	name  string
	age   int
	email string
}

const TABLE_NAME = "users"

const INITIALIZE_USER_TABLE_QUERY = `
	CREATE TABLE IF NOT EXISTS` + TABLE_NAME + ` (
	id serial PRIMARY KEY,
	name text NOT NULL,
	age integer
	email integer)`

func InsertUser(db *sql.DB, user User) (err error) {
	sqlStatement := `
	INSERT INTO $1 (name, age, email)
	VALUES ($2, $3, $4)`
	_, err = db.Exec(sqlStatement, TABLE_NAME, user.name, user.age, user.email)
	return err
}

func DeleteUser(db *sql.DB, user User) (err error) {
	sqlStatement := `
	DELETE FROM $1 WHERE
	name = $1 AND
	age = $1 AND
	email = $3`
	_, err = db.Exec(sqlStatement, TABLE_NAME, user.name, user.age, user.email)
	return err
}
