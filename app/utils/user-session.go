package utils

import . "session-service-v2/app/model"

func FilterSessions(userSessionsData []SessionData, client, fingerPrint string) []SessionData {
	filtered := []SessionData{}
	for i := range userSessionsData {
		if (client != "" && userSessionsData[i].Client == client) && fingerPrint != "" && userSessionsData[i].Fingerprint == fingerPrint {
			filtered = append(filtered, userSessionsData[i])
		}
	}

	if len(filtered) != 0 {
		return filtered
	} else {
		return nil
	}
}

func DeleteFirstClient(userSessionsData []SessionData, client string) []SessionData {
	for i := range userSessionsData {
		if filterClient(userSessionsData[i], client) {
			return append(userSessionsData[:i], userSessionsData[i+1:]...)
		}
	}
	return userSessionsData
}

func DeleteFirst(userSessionsData []SessionData, client, fingerPrint string) []SessionData {
	for i := range userSessionsData {
		if filterClient(userSessionsData[i], client) && filterFingerPrint(userSessionsData[i], fingerPrint) {
			return append(userSessionsData[:i], userSessionsData[i+1:]...)
		}
	}
	return userSessionsData
}

func filterClient(sessionData SessionData, client string) bool {
	return client != "" && sessionData.Client == client
}

func filterFingerPrint(sessionData SessionData, fingerPrint string) bool {
	return fingerPrint != "" && sessionData.Fingerprint == fingerPrint
}
