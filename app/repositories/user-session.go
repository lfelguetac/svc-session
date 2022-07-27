package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	. "session-service-v2/app/model"
	"time"
)

func SetUserSession(userId string, userSession UserSession, ttl string) error {
	us, _ := json.Marshal(userSession)
	_err := redisClient.Set(userId, us, getTtlTime(ttl)).Err()
	if _err != nil {
		fmt.Println(_err)
		return errors.New(_err.Error())
	}
	return nil
}

func GetUserSessions(userId string) (*UserSession, error) {
	result, _err := redisClient.Get(userId).Bytes()
	var userSession UserSession
	json.Unmarshal(result, &userSession)
	if _err != nil {
		fmt.Println(_err)
		return nil, errors.New("id not found")
	}
	return &userSession, nil
}

func DeleteUserSessions(userId string) error {
	_err := redisClient.Del(userId).Err()
	if _err != nil {
		fmt.Println(_err)
		return errors.New(_err.Error())
	}
	return nil
}

func getTtlTime(ttl string) time.Duration {
	ttlHour, _err := time.ParseDuration(ttl)
	if _err != nil {
		return time.Hour
	}
	return ttlHour
}
