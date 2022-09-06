package model

type UserSession struct {
	Sessions []SessionData
}

type SessionRequest struct {
	UserId string      `json:"id" binding:"required"`
	Client string      `json:"client" binding:"required"`
	Ttl    string      `json:"ttl" binding:"required"`
	Data   SessionData `json:"data" binding:"required"`
}

type SessionData struct {
	Token        string `json:"token" binding:"required"`
	RefreshToken string `json:"refreshToken"`
	Fingerprint  string `json:"fingerprint"`
	CoreId       string `json:"core_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Country      string `json:"country"`
	Client       string `json:"client"`
	Ttl          string `json:"ttl"`
}
