package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type apiConfig struct {
	Port uint16
}

type databaseConfig struct {
	Host     string
	Port     uint16
	Database string
	User     string
}

var ApiConfig apiConfig
var DatabaseConfig databaseConfig

func Setup() {
	err := godotenv.Load(".env")
	if err == nil {
		ApiConfig = apiConfig{
			Port: uint16(parseInt(os.Getenv("API_PORT"), 5001)),
		}

		DatabaseConfig = databaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     uint16(parseInt(os.Getenv("DB_POST"), 5432)),
			Database: os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
		}
	}
}

func parseInt(value string, defaultValue int) int {
	vInt, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	} else {
		return vInt
	}
}
