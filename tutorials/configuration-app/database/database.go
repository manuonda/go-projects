package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/manuonda/go-projects/tutorials/configuration-app/configuration"
)

/*

func CreateConnection(ctx context.Context) *sql.DB {

	host, ok := os.LookupEnv("DB_HOST")

	if !ok {
		panic("DB_HOST not set")
	}

	user, ok := os.LookupEnv("DB_USER")
	if !ok {
		panic("DB_USER not set")
	}

	password, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		panic("DB_PASSWORD not set")
	}

	port, ok := os.LookupEnv("DB_PORT")
	if !ok {
		panic("DB_PORT not set")
	}

	database, ok := os.LookupEnv("DB_DATABASE")
	if !ok {
		panic("DB_DATABASE not set")
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
	fmt.Println(connectionString)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}
	return db
}
*/

func CreateConnection(ctx context.Context, config *configuration.Configuration) *sql.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.DB.User,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name)
	fmt.Println(connectionString)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}
	return db

}
