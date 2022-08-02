package repositories_test

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"

	. "session-service-v2/app/model"
	"session-service-v2/app/repositories"
)

func TestSetUserSession(t *testing.T) {

	s := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
	repositories.SetClient(client)

	// TODO: create complete session object and validate fields
	session := SessionData{
		Token: "123",
		RefreshToken: "123"}

	userSession := UserSession{
		Sessions: []SessionData{session},
	}

	repositories.SetUserSession("42", userSession, "1000")

	_, err := s.Get("42")
	if err != nil {
		t.Error("Session is not saved")
	}
}

func TestGetUserSessions(t *testing.T) {

	s := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
	repositories.SetClient(client)

	s.Set("42", "bla")
	// TODO: set entire session object

	_, err := repositories.GetUserSessions("42")
	if err != nil {
		t.Error("Session is not saved")
	}
}

func TestDeleteUserSession(t *testing.T) {

	s := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
	repositories.SetClient(client)

	s.Set("42", "bla")
	// TODO: set entire session object

	err := repositories.DeleteUserSessions("42")
	if err != nil {
		t.Error("Session is not saved")
	}

	_, err = s.Get("42")
	if err == nil {
		t.Error("Session is deleted")
	}
}