package routers

import (
	routers "session-service-v2/app/routers/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/users", routers.CreateUserSession)
	r.GET("/users/:id", routers.GetUserSession)
	r.DELETE("/users/:id", routers.DeleteUserSession)

	return r
}
