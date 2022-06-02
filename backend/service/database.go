package service

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

func DbConnect() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file")
	}
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbPortInt, err := strconv.ParseInt(dbPort, 0, 64)
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	postgresInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPortInt, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", postgresInfo)
	if err != nil {
		panic(err)
	}

	start := time.Now()
	for db.Ping() != nil {
		if start.After(start.Add(10 * time.Second)) {
			fmt.Println("failed to connect after 10 secs.")
			break
		}
	}
	return db
}
