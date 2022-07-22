package repository

import (
	"errors"
	"fmt"
	"session-service-v2/app/model"
)

func CreateUserSession(userSession model.UserSession) error {
	_err := redisClient.Set("asd", userSession, 0).Err()
	if _err != nil {
		fmt.Println(_err)
		return errors.New(_err.Error())
	}
	return nil
}

func GetUserSession(id string) (*model.UserSession, error) {
	_, _err := redisClient.Get(id).Result()
	if _err != nil {
		fmt.Println(_err)
		return nil, errors.New("id not found")
	}

	//TODO convert userSession to struct type and return
	return nil, nil
}
