package services

import (
	"session-service-v2/app/model"
	. "session-service-v2/app/model"
	"session-service-v2/app/repositories"
	. "session-service-v2/app/utils"
)

type UserService interface {
	CreateUserSession(userId, client string, session SessionData, ttl string) error
	GetUserSessions(userId string) (*[]SessionData, error)
	GetUserSession(userId, client string, fingerPrint string) (*SessionData, error)
	DeleteUserSession(userId, client string, fingerPrint string) error
	DeleteUserSessions(userId string) error
}

type service struct {
	repo repositories.UserSessionRepository
}

var multiDevice bool

func NewUserSSService(repository repositories.UserSessionRepository, md bool) UserService {
	multiDevice = md
	return &service{repo: repository}
}

func (svc *service) CreateUserSession(userId string, client string, session model.SessionData, ttl string) error {
	userSession, _err := svc.repo.GetUserSessions(userId)

	// TODO: check only KEY not found error
	if _err != nil {
		userSession := UserSession{
			Sessions: []SessionData{session},
		}
		_err = svc.repo.SetUserSession(userId, userSession, ttl)
	} else {
		if multiDevice {
			sessions := DeleteFirstClient(userSession.Sessions, client)
			userSession.Sessions = append([]SessionData{session}, sessions...)
			_err = svc.repo.SetUserSession(userId, *userSession, ttl)
		} else {
			userSession.Sessions = []SessionData{session}
			_err = svc.repo.SetUserSession(userId, *userSession, ttl)
		}
	}

	if _err != nil {
		return _err
	}
	return nil
}

func (svc *service) GetUserSessions(userId string) (*[]model.SessionData, error) {
	userSession, _err := svc.repo.GetUserSessions(userId)
	if _err != nil {

		return nil, _err
	}
	return &userSession.Sessions, nil
}

// func init() {
// 	multiDevice = GetBoolEnv("MULTIDEVICE_ENABLED", false)
// 	fmt.Println("Multidevice enabled: " + strconv.FormatBool(multiDevice))
// }

func (svc *service) GetUserSession(userId, client, fingerPrint string) (*SessionData, error) {
	userSession, _err := svc.repo.GetUserSessions(userId)
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

func (svc *service) DeleteUserSessions(userId string) error {
	_err := svc.repo.DeleteUserSessions(userId)
	if _err != nil {
		return _err
	}
	return nil
}

func (svc *service) DeleteUserSession(userId, client, fingerPrint string) error {

	userSession, _err := svc.repo.GetUserSessions(userId)
	if _err != nil {
		return _err
	}

	sessions := DeleteFirst(userSession.Sessions, client, fingerPrint)
	if len(sessions) != 0 {
		userSession := UserSession{
			Sessions: sessions,
		}
		//TODO ver que hacer con el TTL cuando se borra una session
		err := svc.repo.SetUserSession(userId, userSession, "1h")
		if err != nil {
			return err
		}
	} else {
		return svc.repo.DeleteUserSessions(userId)
	}
	return nil
}
