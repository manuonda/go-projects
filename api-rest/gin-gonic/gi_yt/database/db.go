package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func Connect() {

	dbError := godotenv.Load()

	if dbError != nil {
		log.Fatal("Error loading .env file %v", dbError.Error())

	}

	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dbFormat := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	fmt.Println("connection url : ", dbFormat)

	DB, err := sql.Open("mysql", dbFormat)
	if err != nil {
		fmt.Println("Fatal error : ", err.Error())
	}

	fmt.Printf("connection success")

	defer DB.Close()

}
