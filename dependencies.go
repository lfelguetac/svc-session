package main

import (
	"session-service-v2/app/delivery/controller"
	"session-service-v2/app/repositories"
	"session-service-v2/app/services"
	"session-service-v2/app/utils"

	"github.com/go-redis/redis/v8"
)

// add & export dependencies between components HERE
func GetDependencies(db *redis.Client) controller.UserSessionController {
	md := utils.GetBoolEnv("MULTIDEVICE_ENABLED", false)

	userRepository := repositories.NewUsersRepository(db)
	userService := services.NewUserService(userRepository, md)
	userController := controller.NewUserSessionController(userService)

	return userController

}
