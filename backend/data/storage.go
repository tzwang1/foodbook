package storage

import (
	"os"
)

var host string = "localhost"
var port int = 5432
var user string = os.Getenv("STORAGE_USER")
var password string = os.Getenv("STORAGE_PASSWORD")
var dbname string = os.Getenv("STORAGE_DB_NAME")

