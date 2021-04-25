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
	VALUES ($2, $3, $4);`
	_, err = db.Exec(sqlStatement, TABLE_NAME, user.name, user.age, user.email)
	return err
}

func UpdateUser(db *sql.DB, user User) (err error) {
	sqlStatement := `
	UPDATE $1 SET
	name = $1,
	age = $2,
	email = $3,
	WHERE id = $4;`
	_, err = db.Exec(sqlStatement, TABLE_NAME, user.name, user.age, user.email, user.id)
	return err
}

func DeleteUser(db *sql.DB, user User) (err error) {
	sqlStatement := `
	DELETE FROM $1 WHERE
	name = $1 AND
	age = $1 AND
	email = $3;`
	_, err = db.Exec(sqlStatement, TABLE_NAME, user.name, user.age, user.email)
	return err
}

func GetUser(db *sql.DB, email string) (User, error) {
	sqlStatement := `
	SELECT * FROM $1 WHERE
	email = $1;`
	row := db.QueryRow(sqlStatement, TABLE_NAME, email)
	var user User
	switch err := row.Scan(&user.id, &user.name, &user.age); err {
	case nil:
		return user, nil
	default:
		return User{}, err
	}
}
