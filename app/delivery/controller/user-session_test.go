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
		"Client": "client123",
		"Ttl":    "ttl123",
		"Data": model.SessionData{
			Token:        "token_1110110010111",
			RefreshToken: "rtokeen_1100011001",
			Fingerprint:  "finger123",
			CoreId:       "coreId123",
			FirstName:    "felipe",
			LastName:     "elgueta",
			Country:      "colbun",
			Client:       "client123",
			Ttl:          "ttl123",
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

	ctrl, mr, reqBody, keyID := BeforeTest(t)

	payload, _ := json.Marshal(reqBody)
	testContextPost, _ := gin.CreateTestContext(httptest.NewRecorder())
	testContextPost.Request, _ = http.NewRequest("POST", "/user", bytes.NewBuffer(payload))
	ctrl.CreateUserSession(testContextPost)

	str, _ := mr.Get(keyID)
	t.Logf("str::: %s", str)

	testContextGet, _ := gin.CreateTestContext(httptest.NewRecorder())
	testContextGet.Request, _ = http.NewRequest("GET", "/user/pepe123", nil)
	ctrl.GetUserSessions(testContextGet)

	expected := 200
	got := testContextGet.Writer.Status()

	t.Logf("keyID %s", keyID)

	if expected != got {
		t.Errorf("expected 200 but got %d", got)
	}

}

func TestGetUserSession(t *testing.T) {

	ctrl, mr, _, keyID := BeforeTest(t)

	reqBody := gin.H{
		"id":     "pepe123",
		"client": "client123",
		"ttl":    "ttl123",
		"data": gin.H{
			"token":        "token_1110110010111",
			"refreshToken": "rtokeen_1100011001",
			"fingerprint":  "finger123",
			"core_Id":      "coreId123",
			"first_name":   "felipe",
			"last_name":    "elgueta",
			"country":      "colbun",
			"client":       "client123",
			"ttl":          "ttl123",
		},
	}

	payload, _ := json.Marshal(reqBody)
	testContextPost, _ := gin.CreateTestContext(httptest.NewRecorder())
	testContextPost.Request, _ = http.NewRequest("POST", "/user", bytes.NewBuffer(payload))
	ctrl.CreateUserSession(testContextPost)

	//
	// 		userRoutes.GET("/:userId/:client/:fingerPrint", handler.userCtrl.GetUserSession)
	w := httptest.NewRecorder()
	testContext, _ := gin.CreateTestContext(w)
	testContext.Request, _ = http.NewRequest("GET", "/user/pepe123/client123/finger123", nil)
	ctrl.GetUserSession(testContext)

	data, _ := mr.Get(keyID)

	t.Logf("datadatadata %s", data)
	t.Logf("code code %d", w.Code)

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
