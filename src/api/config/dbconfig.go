package app

import (
	"go-ginapp/model"
	"log"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
}

var router = gin.Default()

//StartApp is ...
func StartApp() {
	dbdriver := os.Getenv("DB_DRIVER")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_NAME")
	dbport := os.Getenv("DB_PORT")

	_, err := model.Model.Initialize(dbdriver, username, password, dbport, host, database)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	route()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}
	log.Fatal(router.Run(":" + port))
}
