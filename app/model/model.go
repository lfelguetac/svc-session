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

/*
Request actual POST
{
   "id":"", ==> customerId
   "client":"", ==> Enum? FPAY tiene de valor en los logs.
   "clientKey":"", ==> fingerprint
   "ttl":0,
   "data":{
      "core_id":"",
      "first_name":"",
      "last_name":"",
      "fingerprint":"",
      "token":"",
      "refreshToken":"",
      "country":""
   }
}
*/

/*

GET actual
Req /user/${customerId}/${client}/${fingerPrint}
Response
{
   "core_id":"16289294",
   "first_name":"*-*-*-*-*",
   "last_name":"*-*-*-*-*",
   "fingerprint":"eBI0r6GmTOaAUw7198NHcP:APA91bFAgkDYx9",
   "token":"eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9", ==> OAUTH MANAGER USA SOLO ESTE PARA VERIFICAR LA SESSION.
   "refreshToken":"eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.",
   "country":"cl"
}
*/
