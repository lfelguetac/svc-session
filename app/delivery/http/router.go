package http

import (
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

	userRoutes := routes.Group("user")
	{
		userRoutes.POST("/", handler.userCtrl.CreateUserSession)
		userRoutes.GET("/:userId", handler.userCtrl.GetUserSessions)
		userRoutes.GET("/:userId/:client/:fingerPrint", handler.userCtrl.GetUserSession)
		userRoutes.DELETE("/:userId", handler.userCtrl.DeleteUserSessions)
		userRoutes.DELETE("/:userId/:client/:fingerPrint", handler.userCtrl.DeleteUserSession)
	}

}
