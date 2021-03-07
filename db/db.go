package db

import (
	"database/sql"
	"fmt"
	"github.com/3nt3/homework/logging"
	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = 5432
	user     = "homework"
	password = "jQjZKKftp5pXs4f23c2APfobctMSjPRFX8h3W2q69GgfixWBeXdYxXhfaxePKqSi"
	dbname   = "homework"
)

var database *sql.DB

func InitDatabase() error {
	logging.InfoLogger.Printf("connecting to database...\n")
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	foo, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return err
	}
	database = foo

	// close at the end
	defer func() { _ = database.Close() }()

	err = database.Ping()
	if err != nil {
		return err
	}

	logging.InfoLogger.Printf("connection to database successful\n")


	logging.InfoLogger.Printf("creating tables...\n")
	err = initializeTables()
	if err != nil {
		return err
	}


	return nil
}

func initializeTables() error {
	_, err := database.Exec("CREATE TABLE IF NOT EXISTS users (id text PRIMARY KEY UNIQUE, username text UNIQUE, email text UNIQUE, password_hash text, created timestamp, permission int)")
	if err != nil {
		return err
	}
	return nil
}
