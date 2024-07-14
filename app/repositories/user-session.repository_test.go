package repositories_test

import (
	"session-service-v2/app/model"
	"session-service-v2/app/repositories"
	"testing"

	"github.com/alicebob/miniredis/v2"
	redis "github.com/go-redis/redis/v8"
)

func BeforeTest(t *testing.T) (repositories.UserSessionRepository, *miniredis.Miniredis) {
	s := miniredis.RunT(t)
	// miniredis
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
	// defer s.Close()
	userRepo := repositories.NewUsersRepository(client)
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

	t.Run("test err handler", func(t *testing.T) {
		expectedErr := "mock-error"
		s.SetError(expectedErr)
		err := userRepo.SetUserSession(userId, userSession, "t")
		t.Log("errerrerrerrerrerrerr; ", err)
		if err.Error() != expectedErr {
			t.Errorf("ERROR expected %s but got %s", expectedErr, err.Error())
		}
	})

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

	t.Run("test err handler", func(t *testing.T) {
		expected := "id not found"
		s.SetError("mock error")
		userRepo.SetUserSession(userId, userSession, "t")
		_, err := userRepo.GetUserSessions(userId)
		if err.Error() != expected {
			t.Errorf("ERROR expected %s but got %s", expected, err.Error())
		}

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

	t.Run("test err handler", func(t *testing.T) {
		expectedErr := "mock-error-delete"
		s.SetError(expectedErr)
		err := userRepo.DeleteUserSessions(userId)
		if err.Error() != expectedErr {
			t.Errorf("ERROR expected %s but got %s", expectedErr, err.Error())
		}
	})

}
