package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToMySQL() (*sql.DB, error) {
	// Configuración de la conexión
	user := "root"
	password := "root"
	host := "localhost"
	port := 3306
	dbname := "db"

	// Construcción del DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname)

	// Apertura de la conexión
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Verificación de la conexión
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
