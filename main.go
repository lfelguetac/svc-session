package main

import (
	"fmt"
	"log"
	"os"
	"session-service-v2/app/config"
	"session-service-v2/app/delivery/http"
	"session-service-v2/app/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	appEnv := os.Getenv("APP_ENV")
	if appEnv != "prod" {
		_err := godotenv.Load()
		if _err != nil {
			fmt.Println("Error loading .env file" + _err.Error())
		}
	}

	db := config.SetupDBConnection()
	defer config.CloseDBConnection(db)

	app := gin.Default()

	port := utils.GetStringEnv("APP_PORT", "8080")

	userController := GetDependencies(db)

	http.NewAppHandler(app, userController)

	log.Printf("Server stopped, err: %v", app.Run(":"+port))

}
