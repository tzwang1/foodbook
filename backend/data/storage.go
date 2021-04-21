package data

import (
	"database/sql"
	"fmt"
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

func Connect() {
	connected := false
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	for i := 0; i < 10; i++ {
		db, err := sql.Open("postgres", psqlInfo)
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

	}
	if !connected {
		log.Panic("Unable to connect to database.")
	}
	log.Println("Successfully connected!")
}
