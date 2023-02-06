package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = ""
const DB_NAME = "todos"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

var Db *gorm.DB
var err error

func InitDb() (*gorm.DB, error) {
	Db, err := connectDB()
	return Db, err
}

func connectDB() (*gorm.DB, error) {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil, err
	}
	fmt.Println("Connection database realizado")

	return db, err
}
