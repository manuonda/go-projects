package main

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/manuonda/go-projects/tutorials/configuration-app/database"
)

func main() {

	var id int
	var err error

	ctx := context.Background()

	err = godotenv.Load(".env")

	if err != nil {
		panic(err)
	}
	//create connection
	db := database.CreateConnection(ctx)
	defer db.Close()

	// test query
	err = db.QueryRowContext(ctx, "SELECT id from todos").Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Println(id)
	fmt.Println("All good")

}
