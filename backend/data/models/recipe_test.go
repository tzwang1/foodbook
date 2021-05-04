package models

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func newMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

func TestInsertRecipe(t *testing.T) {
	db, mock := newMock()

	recipe := Recipe{Name: "testName", Rating: 0}
	mock.ExpectExec("INSERT INTO recipes").WithArgs(recipe.Name, recipe.Rating).WillReturnResult(sqlmock.NewResult(1, 1))

	err := InsertRecipe(db, recipe)
	assert.NoError(t, err)
}

func TestUpdateRecipe(t *testing.T) {
	db, mock := newMock()
	recipe := Recipe{Id: "testid", Name: "testName", Rating: 1}
	mock.ExpectExec("UPDATE recipes").WithArgs(recipe.Name, recipe.Rating, recipe.Id).WillReturnResult(sqlmock.NewResult(1, 1))

	err := UpdateRecipe(db, recipe)
	assert.NoError(t, err)
}

func TestDeleteRecipe(t *testing.T) {
	db, mock := newMock()
	mock.ExpectExec("DELETE FROM recipes").WithArgs("testid").WillReturnResult(sqlmock.NewResult(1, 1))

	err := DeleteRecipe(db, "testid")
	assert.NoError(t, err)
}

func TestGetRecipe(t *testing.T) {
	db, mock := newMock()
	rows := sqlmock.NewRows([]string{"id", "name", "rating"}).AddRow("testid", "testName", 1)
	mock.ExpectQuery("SELECT \\* FROM recipes WHERE id = \\$1").WithArgs("testid").WillReturnRows(rows)

	recipe, err := GetRecipe(db, "testid")
	assert.NoError(t, err)
	assert.NotNil(t, recipe)
}
