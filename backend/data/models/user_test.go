package models

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var testUser = &User{
	id:    "1",
	name:  "testName",
	age:   5,
	email: "testemail@mail.com"}

func newMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

func TestInsertUser(t *testing.T) {
	db, mock := newMock()

	user := User{name: "testName", age: 1, email: "test@mail.com"}
	mock.ExpectExec("INSERT INTO").WithArgs(USER_TABLE_NAME, user.name, user.age, user.email).WillReturnResult(sqlmock.NewResult(1, 1))

	err := InsertUser(db, user)
	assert.NoError(t, err)
}

func TestUpdateUser(t *testing.T) {
	db, mock := newMock()
	user := User{id: "testid", name: "testName", age: 1, email: "test@mail.com"}
	mock.ExpectExec("UPDATE").WithArgs(USER_TABLE_NAME, user.name, user.age, user.email, user.id).WillReturnResult(sqlmock.NewResult(1, 1))

	err := UpdateUser(db, user)
	assert.NoError(t, err)
}
