package data

import (
	"database/sql"
	"fmt"
	"foodbook/backend/data/models"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var host string = "host.docker.internal"
var port int = 5432
var user string = os.Getenv("POSTGRES_USER")
var password string = os.Getenv("POSTGRES_PASSWORD")
var dbname string = os.Getenv("POSTGRES_DB")

type Database struct {
	Db *sql.DB
}

var instance *Database

func GetDatabaseSingleton() *Database {
	if instance == nil {
		instance = &Database{Db: InitializeDatabase()}
	}
	return instance
}

func InitializeDatabase() *sql.DB {
	connected := false
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var db *sql.DB
	var err error
	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Println("Unable to Open DB: " + err.Error() + " ... Retrying\n")
			time.Sleep(time.Second * 2)
			continue
		}
		if err = db.Ping(); err != nil {
			log.Println("Unable to Ping DB: " + err.Error() + " ... Retrying\n")
			time.Sleep(time.Second * 2)
			continue
		}
		connected = true
		break
	}
	if !connected {
		log.Panic("Unable to connect to database.")
	}
	log.Println("Successfully connected!")
	log.Println("Initializing Tables...")
	err = initializeTables(db)
	if err != nil {
		log.Fatalf("Failed to initialize database tables with error: %v\n", err)
	}
	return db
}

func initializeTables(db *sql.DB) (err error) {
	initialization_queries := []string{models.INITIALIZE_RECIPE_TABLE_QUERY,
		models.INITIALIZE_INSTRUCTION_TABLE_QUERY}

	for _, initialize_query := range initialization_queries {
		_, err = db.Exec(initialize_query)
		if err != nil {
			return
		}
	}
	// return InitializeRecipes()
	return nil
}
