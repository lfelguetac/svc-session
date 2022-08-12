package main

import (
	"session-service-v2/app/config"
	"session-service-v2/app/delivery/controller"
	"session-service-v2/app/delivery/http"
	"session-service-v2/app/repositories"
	"session-service-v2/app/services"
	"session-service-v2/app/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetupDBConnection()
	// defer config.CloseDBConnection(db)

	r := gin.Default()

	md := utils.GetBoolEnv("MULTIDEVICE_ENABLED", false)

	userRepository := repositories.NewUserSsRepository(db)
	userService := services.NewUserSSService(userRepository, md)
	userController := controller.NewUserSessionController(userService)

	http.NewAppHandler(r, userController)

	r.Run()
}
