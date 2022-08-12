package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"session-service-v2/app/delivery/controller"
	"session-service-v2/app/model"
	"session-service-v2/app/repositories"
	"session-service-v2/app/services"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/gin-gonic/gin"

	// "github.com/go-redis/redis"
	redis "github.com/go-redis/redis/v8"
)

func BeforeTest(t *testing.T) (controller.UserSessionController, *miniredis.Miniredis, gin.H, string) {
	s := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	keyID := "pepe123"

	reqBody := gin.H{
		"Id":     keyID,
		"Client": "client12",
		"Ttl":    "001001101",
		"Data": model.SessionData{
			Token:        "4554545",
			RefreshToken: "6566666",
			Fingerprint:  "finger12",
			CoreId:       "1212",
			FirstName:    "felipe",
			LastName:     "elgueta",
			Country:      "colbun",
			Client:       "111",
			Ttl:          "9874",
		},
	}

	userRepository := repositories.NewUserSsRepository(client)
	userService := services.NewUserSSService(userRepository, true)
	userController := controller.NewUserSessionController(userService)
	return userController, s, reqBody, keyID

}

func TestCreateUserSession(t *testing.T) {

	ctrl, _, reqBody, _ := BeforeTest(t)
	w := httptest.NewRecorder()
	testContext, _ := gin.CreateTestContext(w)

	t.Run("ShouldBindJSON error", func(t *testing.T) {

		payload, _ := json.Marshal(reqBody)
		testContext.Request, _ = http.NewRequest("POST", "/user", bytes.NewBuffer(payload))
		testContext.Request.Header.Set("Content-Type", "application/json")

		ctrl.CreateUserSession(testContext)

		expected := 201
		got := testContext.Writer.Status()

		if expected != got {
			t.Errorf("expected 200 but got %d", got)
		}

	})

}

func TestGetUserSessions(t *testing.T) {

	ctrl, _, reqBody, keyID := BeforeTest(t)
	w := httptest.NewRecorder()
	testContext, _ := gin.CreateTestContext(w)

	payload, _ := json.Marshal(reqBody)
	testContext.Request, _ = http.NewRequest("POST", "/user", bytes.NewBuffer(payload))
	// testContext.Request.Header.Set("Content-Type", "application/json")
	ctrl.CreateUserSession(testContext)
	// bytes, _ := json.Marshal(reqBody)
	// mr.Set(keyID, "bytes")

	ctrl.GetUserSessions(testContext)
	testContext.Request, _ = http.NewRequest("GET", "/user/"+keyID, nil)

	expected := 200
	got := testContext.Writer.Status()

	if expected != got {
		t.Errorf("expected 200 but got %d", got)
	}

}

func TestGetUserSession(t *testing.T) {

	ctrl, _, _, keyID := BeforeTest(t)
	w := httptest.NewRecorder()
	testContext, _ := gin.CreateTestContext(w)

	testContext.Request, _ = http.NewRequest("GET", "/user/"+keyID+"/client12/finger12", nil)
	ctrl.GetUserSession(testContext)

	expected := 200
	got := testContext.Writer.Status()

	if expected != got {
		t.Errorf("expected 200 but got %d", got)
	}
}

func TestDeleteUserSessions(t *testing.T) {

	ctrl, _, _, keyID := BeforeTest(t)
	w := httptest.NewRecorder()
	testContext, _ := gin.CreateTestContext(w)

	// mr.Set(keyID, "anything")
	testContext.Request, _ = http.NewRequest("DELETE", "/user/"+keyID, nil)
	ctrl.DeleteUserSessions(testContext)

	expected := 200
	got := testContext.Writer.Status()

	if expected != got {
		t.Errorf("expected 200 but got %d", got)
	}

}

func TestDeleteUserSession(t *testing.T) {

	ctrl, mr, _, keyID := BeforeTest(t)
	w := httptest.NewRecorder()
	testContext, _ := gin.CreateTestContext(w)

	mr.Set(keyID, "anything")
	testContext.Request, _ = http.NewRequest("DELETE", "/user/"+keyID+"/client12/finger12", nil)
	ctrl.DeleteUserSession(testContext)

	expected := 200
	got := testContext.Writer.Status()

	if expected != got {
		t.Errorf("expected 200 but got %d", got)
	}

}
