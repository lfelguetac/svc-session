package model

type UserSession struct {
	Sessions []Session
}

type Session struct {
	Jwt         string `json:"jwt"`
	Fingerprint string `json:"fingerprint"`
	Ttl         string `json:"ttl"`
}

type SessionRequest struct {
	User        string `json:"user"`
	Jwt         string `json:"jwt"`
	Fingerprint string `json:"fingerprint"`
	Ttl         string `json:"ttl"`
}
