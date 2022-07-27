package routers

import (
	routers "session-service-v2/app/routers/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/user", routers.CreateUserSession)
	r.GET("/user/:userId", routers.GetUserSessions)
	r.GET("/user/:userId/:client/:fingerPrint", routers.GetUserSession)
	r.DELETE("/user/:userId", routers.DeleteUserSessions)
	r.DELETE("/user/:userId/:client/:fingerPrint", routers.DeleteUserSession)

	return r
}
