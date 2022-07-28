package routers

import (
	"net/http"
	"session-service-v2/app/logger"
	. "session-service-v2/app/model"
	"session-service-v2/app/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log *logrus.Entry = logger.GetLogger()

func CreateUserSession(c *gin.Context) {

	log.Info("Trying to create user session")

	req := SessionRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

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

	_err := services.CreateUserSession(userId, client, session, ttl)

	if _err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error creating"})
		return
	}
	c.AbortWithStatus(http.StatusCreated)
}

func GetUserSessions(c *gin.Context) {
	userId := c.Param("userId")

	sessions, _err := services.GetUserSessions(userId)
	if _err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, sessions)
	return
}

func GetUserSessionsForClient(c *gin.Context) {
	userId := c.Param("userId")
	client := c.Param("client")

	sessions, _err := services.GetUserSession(userId, client, "")
	if _err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, sessions)
	return
}

func GetUserSession(c *gin.Context) {
	userId := c.Param("userId")
	client := c.Param("client")
	fingerPrint := c.Param("fingerPrint")

	session, _err := services.GetUserSession(userId, client, fingerPrint)
	if _err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, session)
	return
}

func DeleteUserSessions(c *gin.Context) {
	userId := c.Param("userId")
	_err := services.DeleteUserSessions(userId)
	if _err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.AbortWithStatus(http.StatusOK)
	return
}

func DeleteUserSession(c *gin.Context) {
	userId := c.Param("userId")
	client := c.Param("client")
	fingerPrint := c.Param("fingerPrint")

	_err := services.DeleteUserSession(userId, client, fingerPrint)
	if _err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.AbortWithStatus(http.StatusOK)
	return
}
