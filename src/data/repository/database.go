package database

// import (
// 	"database/sql"
// 	"digitalwallet-service/src/webapp/config"
// 	"fmt"
// 	"log"
// 	"os"

// 	_ "github.com/lib/pq"
// )

// var pgDatabase *sql.DB

// func ExecStatement(query string, args ...any) (int64, error) {
// 	statement, err := pgDatabase.Prepare(query)
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer statement.Close()

// 	result, err := statement.Exec(args...)
// 	if err != nil {
// 		return 0, err
// 	}

// 	rows, err := result.RowsAffected()
// 	if err != nil {
// 		return 0, err
// 	}

// 	if rows == 0 {
// 		return 0, fmt.Errorf("Transactions was not inserted")
// 	}

// 	return rows, nil
// }

// func GetConnection() *sql.DB {
// 	if pgDatabase == nil {
// 		return openConnection()
// 	}

// 	if err := pgDatabase.Ping(); err != nil {
// 		pgDatabase.Close()
// 		return openConnection()
// 	}

// 	return pgDatabase
// }

// func openConnection() *sql.DB {
// 	dbConfig := config.DatabaseConfig
// 	stringConnection := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
// 		dbConfig.Host, dbConfig.Port, dbConfig.Database, dbConfig.User, os.Getenv("DB_PASSWORD"))

// 	database, err := sql.Open("postgres", stringConnection)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if err := database.Ping(); err != nil {
// 		database.Close()
// 		log.Fatal(err)
// 	}

// 	database.SetMaxOpenConns(90)
// 	database.SetConnMaxLifetime(0)

// 	pgDatabase = database

// 	return database
// }
