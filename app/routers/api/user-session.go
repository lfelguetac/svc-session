package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"session-service-v2/app/model"
	"session-service-v2/app/repository"
)

func CreateUserSession(c *gin.Context) {
	req := model.SessionRequest{}
	c.ShouldBindJSON(&req)
	//
	//session := model.Session{
	//	Jwt:         req.Jwt,
	//	Fingerprint: req.Fingerprint,
	//	Ttl:         req.Ttl,
	//}
	//
	//userSession, err := repository.GetUserSession(req.User)
	//
	//if err != nil {
	//	userSession.Sessions: append(userSession.Sessions, session)
	//
	//	err = repository.CreateUserSession(userSession)
	//} else {
	//	userSession := model.UserSession{
	//		Sessions: []model.Session{session},
	//	}
	//	err = repository.CreateUserSession(userSession)
	//}
	//
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"message": "error creating"})
	//	return
	//}
	c.AbortWithStatus(http.StatusCreated)
}

func GetUserSession(c *gin.Context) {
	id := c.Param("id")
	userSession, err := repository.GetUserSession(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(200, userSession)
	return
}
