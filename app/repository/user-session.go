package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	. "session-service-v2/app/model"
)

func SetUserSession(userId string, userSession UserSession, ttl int) error {
	us, _ := json.Marshal(userSession)

	_err := redisClient.Set(userId, us, 0).Err()
	if _err != nil {
		fmt.Println(_err)
		return errors.New(_err.Error())
	}
	return nil
}

func GetUserSession(id string) (*UserSession, error) {
	result, _err := redisClient.Get(id).Bytes()
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