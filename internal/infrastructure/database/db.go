package database

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func ConnectToDatabase() (*gorm.DB, error) {
	dbConnection := os.Getenv("DB_CONNECTION")

	switch dbConnection {
	case "mysql":
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)

		conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}

		return conn, nil
	case "postgres":
		dsn := fmt.Sprintf(
			"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
		)

		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}

		return conn, nil
	default:
		return nil, errors.New("Cannot connect to database.")
	}
}
