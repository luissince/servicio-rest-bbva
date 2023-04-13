package database

import (
	"database/sql"
	"fmt"

	_ "github.com/microsoft/go-mssqldb"
)

// CreateConnection -> Crea la conexion a la base de datos
func CreateConnection() (*sql.DB, error) {

	// Creamos la cadena de conexión
	server := "10.1.4.9"
	port := 1435
	user := "csedanoq"
	password := "123456"
	database := "SGA"

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)

	// Creamos una instancia de sql.DB
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(5)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
