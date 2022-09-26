package http

import (
	"net/http"
	"os"
	"session-service-v2/app/delivery/controller"

	"github.com/gin-gonic/gin"
)

type appHandler struct {
	userCtrl controller.UserSessionController
}

func NewAppHandler(routes *gin.Engine, userCtrl controller.UserSessionController) {

	handler := &appHandler{
		userCtrl: userCtrl,
	}

	routes.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "App listening on port "+os.Getenv("APP_PORT"))
	})

	userRoutes := routes.Group("session/user")
	{
		userRoutes.POST("/", handler.userCtrl.CreateUserSession)
		userRoutes.POST("", handler.userCtrl.CreateUserSession)
		userRoutes.GET("/:userId", handler.userCtrl.GetUserSessions)
		userRoutes.GET("/:userId/:client/:fingerPrint", handler.userCtrl.GetUserSession)
		userRoutes.DELETE("/:userId", handler.userCtrl.DeleteUserSessions)
		userRoutes.DELETE("/:userId/:client/:fingerPrint", handler.userCtrl.DeleteUserSession)
	}

}
