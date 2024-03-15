package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type appConfig struct {
	ENVIRONMENT string
	PORT        string

	DB_CONNECTION string
	DB_HOST       string
	DB_PORT       int32
	DB_NAME       string
	DB_USERNAME   string
	DB_PASSWORD   string
}

var AppConfig appConfig

func LoadConfig() error {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("💣💥 Error parsing DB_PORT")
		return err
	}

	dbPortStr := os.Getenv("DB_PORT")
	dbPort, err := strconv.ParseInt(dbPortStr, 10, 32)
	if err != nil {
		log.Fatal("💣💥 Error parsing DB_PORT")
		return err
	}

	AppConfig.ENVIRONMENT = os.Getenv("ENVIRONMENT")
	AppConfig.PORT = os.Getenv("PORT")

	AppConfig.DB_CONNECTION = os.Getenv("DB_CONNECTION")
	AppConfig.DB_HOST = os.Getenv("DB_HOST")
	AppConfig.DB_PORT = int32(dbPort)
	AppConfig.DB_NAME = os.Getenv("DB_NAME")
	AppConfig.DB_USERNAME = os.Getenv("DB_USERNAME")
	AppConfig.DB_PASSWORD = os.Getenv("DB_PASSWORD")

	return nil
}
