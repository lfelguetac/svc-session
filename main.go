package main

import (
	"fmt"
	"session-service-v2/app/routers"
	. "session-service-v2/app/utils"
)

func main() {
	fmt.Println("Starting API Sessions")

	// Init Gin
	r := routers.InitRouter()


	port := GetStringEnv("APP_PORT", "8080")

	// Run server
	r.Run(":" + port)
}
