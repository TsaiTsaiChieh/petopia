package main

import (
	"fmt"
	"log"
	"net/http"
	"petopia-server/jobs"
	models "petopia-server/models"
	routes "petopia-server/router"
	services "petopia-server/services"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {
	initializeEnvironment()
	initDatabase()
	scheduler()
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

func scheduler() {
	c := cron.New()

	c.AddFunc("@every 30s", func() {
		fmt.Println("Every ten second")
		jobs.DailyTask()
	})
	c.Start()

	fmt.Println("Program has started. Waiting for scheduled tasks to execute...")

}
