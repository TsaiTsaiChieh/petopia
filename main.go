package main

import (
	"log"
	"net/http"
	routes "petopia-server/router"
	services "petopia-server/services"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	db, err := services.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()
	router := routes.NewRouter()
	http.ListenAndServe(":3000", router)
}
