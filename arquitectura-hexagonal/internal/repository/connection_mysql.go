package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToMySQL() (*sql.DB, error) {
	// Configuración de la conexión
	user := "tu_usuario"
	password := "tu_contraseña"
	host := "localhost"
	port := 3306
	dbname := "tu_base_de_datos"

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
