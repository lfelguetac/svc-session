package services

import (
	"fmt"
	. "session-service-v2/app/model"
	"session-service-v2/app/repositories"
	. "session-service-v2/app/utils"

	"strconv"
)

var multiDevice bool

func init() {
	multiDevice = GetBoolEnv("MULTIDEVICE_ENABLED", false)
	fmt.Println("Multidevice enabled: " + strconv.FormatBool(multiDevice))
}

func CreateUserSession(userId, client string, session SessionData, ttl string) error {
	userSession, _err := repositories.GetUserSessions(userId)

	// TODO: check only KEY not found error
	if _err != nil {
		userSession := UserSession{
			Sessions: []SessionData{session},
		}
		_err = repositories.SetUserSession(userId, userSession, ttl)
	} else {
		if multiDevice {
			sessions := DeleteFirstClient(userSession.Sessions, client)
			userSession.Sessions = append([]SessionData{session}, sessions...)
			_err = repositories.SetUserSession(userId, *userSession, ttl)
		} else {
			userSession.Sessions = []SessionData{session}
			_err = repositories.SetUserSession(userId, *userSession, ttl)
		}
	}

	if _err != nil {
		return _err
	}
	return nil
}

func GetUserSessions(userId string) (*[]SessionData, error) {
	userSession, _err := repositories.GetUserSessions(userId)
	if _err != nil {
		return nil, _err
	}
	return &userSession.Sessions, nil
}

func GetUserSession(userId, client, fingerPrint string) (*SessionData, error) {
	userSession, _err := repositories.GetUserSessions(userId)
	if _err != nil {
		return nil, _err
	}

	if multiDevice {
		sessions := FilterSessions(userSession.Sessions, client, fingerPrint)
		if sessions != nil {
			return &sessions[0], nil
		} else {
			return nil, nil
		}
	} else {
		return &userSession.Sessions[0], nil
	}
}

func DeleteUserSessions(userId string) error {
	_err := repositories.DeleteUserSessions(userId)
	if _err != nil {
		return _err
	}
	return nil
}

func DeleteUserSession(userId, client, fingerPrint string) error {
	userSession, _err := repositories.GetUserSessions(userId)
	if _err != nil {
		return _err
	}
	sessions := DeleteFirst(userSession.Sessions, client, fingerPrint)
	if len(sessions) != 0 {
		userSession := UserSession{
			Sessions: sessions,
		}
		//TODO ver que hacer con el TTL cuando se borra una session
		err := repositories.SetUserSession(userId, userSession, "1h")
		if err != nil {
			return err
		}
	} else {
		return DeleteUserSessions(userId)
	}
	return nil
}
