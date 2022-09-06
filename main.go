package main

import (
	"session-service-v2/app/config"
	"session-service-v2/app/delivery/http"
	"session-service-v2/app/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetupDBConnection()
	defer config.CloseDBConnection(db)

	gin := gin.Default()

	port := utils.GetStringEnv("APP_PORT", "8080")

	userController := GetDependencies(db)

	http.NewAppHandler(gin, userController)

	gin.Run(":" + port)
}
