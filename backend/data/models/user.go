package models

import (
	"database/sql"
)

type User struct {
	Id    string
	Name  string
	Age   int
	Email string
}

const USER_TABLE_NAME = "users"

const INITIALIZE_USER_TABLE_QUERY = `
	CREATE TABLE IF NOT EXISTS ` + USER_TABLE_NAME + ` (
	id serial PRIMARY KEY,
	name text NOT NULL,
	age integer,
	email text UNIQUE)`

func InsertUser(db *sql.DB, user User) (err error) {
	sqlStatement := `INSERT INTO users(name, age, email) VALUES ($1, $2, $3);`
	_, err = db.Exec(sqlStatement, user.Name, user.Age, user.Email)
	return err
}

func UpdateUser(db *sql.DB, user User) (err error) {
	sqlStatement := `
	UPDATE $1 SET
	name = $1,
	age = $2,
	email = $3,
	WHERE id = $4;`
	_, err = db.Exec(sqlStatement, USER_TABLE_NAME, user.Name, user.Age, user.Email, user.Id)
	return err
}

func DeleteUser(db *sql.DB, user User) (err error) {
	sqlStatement := `
	DELETE FROM $1 WHERE
	name = $1 AND
	age = $1 AND
	email = $3;`
	_, err = db.Exec(sqlStatement, USER_TABLE_NAME, user.Name, user.Age, user.Email)
	return err
}

func GetUser(db *sql.DB, email string) (User, error) {
	sqlStatement := `
	SELECT * FROM users WHERE email = $1;`
	row := db.QueryRow(sqlStatement, email)
	var user User
	switch err := row.Scan(&user.Id, &user.Name, &user.Age, &user.Email); err {
	case nil:
		return user, nil
	default:
		return User{}, err
	}
}
