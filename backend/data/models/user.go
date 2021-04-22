package models

import (
	"database/sql"
)

type User struct {
	id   string
	name string
	age  int
}