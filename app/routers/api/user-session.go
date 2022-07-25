package routers

import (
	"fmt"
	"net/http"
	. "session-service-v2/app/model"
	"session-service-v2/app/repository"
	. "session-service-v2/app/utils"

	"github.com/gin-gonic/gin"

	"strconv"
)

var multiDevice bool

func init() {
	multiDevice = GetBoolEnv("MULTIDEVICE_ENABLED", false)
	fmt.Println("Multidevice enabled: " + strconv.FormatBool(multiDevice))
}

func CreateUserSession(c *gin.Context) {
	req := SessionRequest{}
	c.ShouldBindJSON(&req)

	session := Session{
		Jwt:         req.Jwt,
		Fingerprint: req.Fingerprint,
		Ttl:         req.Ttl,
	}

	userSession, err := repository.GetUserSession(req.User)

	if err != nil {
		userSession := UserSession{
			Sessions: []Session{session},
		}
		err = repository.SetUserSession(req.User, userSession, req.Ttl)
	} else {
		if multiDevice {
			userSession.Sessions = append([]Session{session}, userSession.Sessions...)
			err = repository.SetUserSession(req.User, *userSession, req.Ttl)
		} else {
			userSession.Sessions = []Session{session}
			err = repository.SetUserSession(req.User, *userSession, req.Ttl)
		}
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error creating"})
		return
	}
	c.AbortWithStatus(http.StatusCreated)
}

func GetUserSession(c *gin.Context) {
	id := c.Param("id")
	userSession, err := repository.GetUserSession(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, userSession)
	return
}

func DeleteUserSession(c *gin.Context) {
	id := c.Param("id")
	err := repository.DeleteUserSession(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.AbortWithStatus(http.StatusOK)
	return
}
