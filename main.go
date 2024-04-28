package main

import (
	"log"
	"net/http"
	"petopia-server/models"
	routes "petopia-server/router"
	services "petopia-server/services"

	"github.com/joho/godotenv"
)

func main() {
	initializeEnvironment()
	initDatabase()
	startServer()
}

func initializeEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func initDatabase() {
	db, err := services.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	err = db.AutoMigrate(&models.Animal{})
	if err != nil {
		log.Fatal("Error auto migrating tables:", err)
	}
}

func startServer() {
	router := routes.NewRouter()
	http.ListenAndServe(":3000", router)
}
