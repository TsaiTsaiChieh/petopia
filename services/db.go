package services

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	username := os.Getenv("username")
	password := os.Getenv("password")
	hostname := os.Getenv("hostname")
	port := os.Getenv("port")
	database := os.Getenv("database")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	log.Println("Connected to MariaDB!")
	return db, nil
}
