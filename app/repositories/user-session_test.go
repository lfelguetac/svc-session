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