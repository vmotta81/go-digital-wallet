package database

import (
	"database/sql"
	"digitalwallet-service/src/webapp/config"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var pgDatabase *sql.DB

func GetConnection() *sql.DB {
	if pgDatabase == nil {
		return openConnection()
	}

	if err := pgDatabase.Ping(); err != nil {
		pgDatabase.Close()
		return openConnection()
	}

	return pgDatabase
}

func openConnection() *sql.DB {
	dbConfig := config.DatabaseConfig
	stringConnection := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.Database, dbConfig.User, os.Getenv("DB_PASSWORD"))

	println(stringConnection)

	database, err := sql.Open("postgres", stringConnection)
	if err != nil {
		log.Fatal(err)
	}

	if err := database.Ping(); err != nil {
		database.Close()
		log.Fatal(err)
	}

	pgDatabase = database

	return database
}
