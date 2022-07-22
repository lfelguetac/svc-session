package main

import (
	"fmt"
	"session-service-v2/app/routers"
)

func main() {
	fmt.Println("Starting Mock Server")

	// Init Gin
	r := routers.InitRouter()

	// Run server
	r.Run()
}
