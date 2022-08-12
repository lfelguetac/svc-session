package repositories_test

import (
	"session-service-v2/app/model"
	"session-service-v2/app/repositories"
	"testing"

	"github.com/alicebob/miniredis/v2"
	// "github.com/go-redis/redis"
	redis "github.com/go-redis/redis/v8"
)

func BeforeTest(t *testing.T) (repositories.UserSessionRepository, *miniredis.Miniredis) {
	s := miniredis.RunT(t)
	// miniredis
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
	// defer s.Close()
	userRepo := repositories.NewUserSsRepository(client)
	return userRepo, s
}

func TestSetUserSession(t *testing.T) {

	userRepo, s := BeforeTest(t)

	session := model.SessionData{
		Token:        "123",
		RefreshToken: "123"}

	userSession := model.UserSession{
		Sessions: []model.SessionData{session},
	}

	userId := "pepe"
	userRepo.SetUserSession(userId, userSession, "t")

	_, err := s.Get(userId)
	if err != nil {
		t.Error("Session is not saved ", err)
	}

}

func TestGetUserSessions(t *testing.T) {
	userRepo, s := BeforeTest(t)

	userId := "123"
	session := model.SessionData{
		Token:        "123",
		RefreshToken: "123"}

	userSession := model.UserSession{
		Sessions: []model.SessionData{session},
	}

	userRepo.SetUserSession(userId, userSession, "t")
	user, _ := userRepo.GetUserSessions(userId)

	if user == nil {
		t.Errorf("ERROR expected %s but got nil", userSession)
	}

	t.Run("multiDevice  false", func(t *testing.T) {

		// mr := miniredis.RunT(t)
		// client := redis.NewClient(&redis.Options{
		// 	Addr: mr.Addr(),
		// })

		// userSsRepo := repositories.NewUserSsRepository(client)
		// svc := services.NewUserSSService(userSsRepo, false)

		// mr.Set(userID, "anything")
		// err := svc.CreateUserSession(userID, "client", mockSession, "lala")

		// if err != nil {
		// 	t.Errorf("FAIL err: expected nil but got %s ", err)
		// }
		defer s.Close()
	})

}

func TestDeleteUserSessions(t *testing.T) {
	userRepo, s := BeforeTest(t)

	userId := "pepe"

	session := model.SessionData{
		Token:        "123",
		RefreshToken: "123",
	}

	userSession := model.UserSession{
		Sessions: []model.SessionData{session},
	}

	userRepo.SetUserSession(userId, userSession, "t")

	userRepo.DeleteUserSessions(userId)

	_, err := s.Get(userId)
	if err == nil {
		t.Errorf("expected %s but got nil", err)
	}

	// t.Run("err != nil", func(t *testing.T) {

	// })

}
