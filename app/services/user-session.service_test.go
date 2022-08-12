package services_test

import (
	"session-service-v2/app/model"
	"session-service-v2/app/repositories"
	"session-service-v2/app/services"
	"testing"

	"github.com/alicebob/miniredis/v2"
	// "github.com/go-redis/redis"
	redis "github.com/go-redis/redis/v8"
)

func BeforeTest(t *testing.T, md bool) (services.UserService, *miniredis.Miniredis, model.SessionData) {

	s := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
	userSsRepo := repositories.NewUserSsRepository(client)

	mockSession := model.SessionData{
		Token:        "4554545",
		RefreshToken: "6566666",
		Fingerprint:  "121212",
		CoreId:       "1212",
		FirstName:    "felipe",
		LastName:     "elgueta",
		Country:      "colbun",
		Client:       "111",
		Ttl:          "9874",
	}

	return services.NewUserSSService(userSsRepo, md), s, mockSession
}

func TestCreateUserSession(t *testing.T) {

	userSvc, mr, mockSession := BeforeTest(t, true)
	userID := "pepito123"

	err := userSvc.CreateUserSession(userID, "client", mockSession, "lala")

	if err != nil {
		t.Errorf("FAIL err: expected nil but got %s ", err)
	}

	t.Run("err = nil", func(t *testing.T) {

		mr.Set(userID, "anything")
		err := userSvc.CreateUserSession(userID, "client", mockSession, "lala")

		if err != nil {
			t.Errorf("FAIL err: expected nil but got %s ", err)
		}

	})

	t.Run("multiDevice  false", func(t *testing.T) {

		mr := miniredis.RunT(t)
		client := redis.NewClient(&redis.Options{
			Addr: mr.Addr(),
		})

		userSsRepo := repositories.NewUserSsRepository(client)
		svc := services.NewUserSSService(userSsRepo, false)

		mr.Set(userID, "anything")
		err := svc.CreateUserSession(userID, "client", mockSession, "lala")

		if err != nil {
			t.Errorf("FAIL err: expected nil but got %s ", err)
		}

	})

}

func TestGetUserSession(t *testing.T) {

	userSvc, _, mockSession := BeforeTest(t, true)

	keyID := "pepe123"
	userSvc.CreateUserSession(keyID, "client", mockSession, "lala")
	_, err := userSvc.GetUserSession(keyID, "client", "lala")

	if err != nil {
		t.Errorf("FAIL err: expected nil but got %s ", err)
	}

}

func TestGetUserSessions(t *testing.T) {

	userSvc, _, mockSession := BeforeTest(t, true)

	keyID := "pepe123"
	userSvc.CreateUserSession(keyID, "client", mockSession, "lala")
	_, err := userSvc.GetUserSessions(keyID)

	if err != nil {
		t.Errorf("FAIL err: expected nil but got %s ", err)
	}
}

func TestDeleteUserSessions(t *testing.T) {

	userSvc, s, mockSession := BeforeTest(t, true)

	keyID := "pepe123"

	userSvc.CreateUserSession(keyID, "client", mockSession, "lala")

	userSvc.DeleteUserSessions(keyID)

	user, err := s.Get(keyID)

	if err == nil {
		t.Logf("FAIL user: expected key not fund  but got %s ", user)
	}

}

func TestDeleteUserSession(t *testing.T) {

	userSvc, s, mockSession := BeforeTest(t, true)

	keyID := "pepe123"

	userSvc.CreateUserSession(keyID, "client", mockSession, "lala")
	userSvc.DeleteUserSession(keyID, "client", "lala")

	user, err := s.Get(keyID)

	if err == nil {
		t.Errorf("FAIL user: expected key not fund  but got %s ", user)
	}

}
