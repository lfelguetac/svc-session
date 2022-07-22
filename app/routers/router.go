package routers

import (
	"github.com/gin-gonic/gin"
	routers "session-service-v2/app/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/users", routers.CreateUserSession)
	r.GET("/users/:id", routers.GetUserSession)

	return r
}
