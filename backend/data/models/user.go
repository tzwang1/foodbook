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
	email integer UNIQUE)`

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

func GetUser(db *sql.DB, email string) (User, error) {
	sqlStatement := `
	SELECT * FROM $1 WHERE
	email = $1`
	row := db.QueryRow(sqlStatement, TABLE_NAME, email)
	var id string
	var name string
	var age int
	switch err := row.Scan(&id, &name, &age); err {
	case nil:
		return User{id: id, name: name, age: age, email: email}, nil
	default:
		return User{}, err
	}
}
