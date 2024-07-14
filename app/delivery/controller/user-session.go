package controller

import (
	"net/http"
	"session-service-v2/app/logger"
	. "session-service-v2/app/model"
	"session-service-v2/app/services"

	"github.com/gin-gonic/gin"
)

type UserSessionController interface {
	CreateUserSession(c *gin.Context)
	GetUserSessions(c *gin.Context)
	GetUserSession(c *gin.Context)
	DeleteUserSession(c *gin.Context)
	DeleteUserSessions(c *gin.Context)
}

// this struct used to accommodate all the services needed
type controller struct {
	userService services.UserService
}

func NewUserSessionController(userSvc services.UserService) UserSessionController {
	return &controller{
		userService: userSvc,
	}
}

var log *logger.FpayLogger = logger.GetLogger()

func (ctr *controller) CreateUserSession(c *gin.Context) {

	req := SessionRequest{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	log.Info("Trying to create user session", map[string]string{"userId": req.UserId})

	userId, client, ttl := req.UserId, req.Client, req.Ttl

	session := SessionData{
		Token:        req.Data.Token,
		RefreshToken: req.Data.RefreshToken,
		Fingerprint:  req.Data.Fingerprint,
		CoreId:       req.Data.CoreId,
		FirstName:    req.Data.FirstName,
		LastName:     req.Data.LastName,
		Country:      req.Data.Country,
		Client:       req.Client,
		Ttl:          req.Ttl,
	}

	_err := ctr.userService.CreateUserSession(userId, client, session, ttl)

	if _err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error creating"})
		return
	}
	c.AbortWithStatus(http.StatusCreated)
}

func (ctr *controller) GetUserSessions(c *gin.Context) {

	userId := c.Param("userId")

	sessions, _err := ctr.userService.GetUserSessions(userId)
	if _err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, sessions)

}

func (ctr *controller) GetUserSession(c *gin.Context) {
	userId := c.Param("userId")
	client := c.Param("client")
	fingerPrint := c.Param("fingerPrint")

	session, _err := ctr.userService.GetUserSession(userId, client, fingerPrint)
	if _err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, session)
}

func (ctr *controller) DeleteUserSessions(c *gin.Context) {
	userId := c.Param("userId")
	_err := ctr.userService.DeleteUserSessions(userId)
	if _err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.AbortWithStatus(http.StatusOK)
}

func (ctr *controller) DeleteUserSession(c *gin.Context) {
	userId := c.Param("userId")
	client := c.Param("client")
	fingerPrint := c.Param("fingerPrint")

	_err := ctr.userService.DeleteUserSession(userId, client, fingerPrint)
	if _err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.AbortWithStatus(http.StatusOK)
}
