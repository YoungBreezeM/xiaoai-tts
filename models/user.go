package models

import "fmt"

type MiAccount struct {
	User string
	Pwd  string
}

type Session struct {
	ServiceToken string
	UserId       string
	DeviceId     string
	SerialNumber string
}

func (s *Session) GetCookie() string {

	cookie := fmt.Sprintf("userId=%s;serviceToken=%s", s.UserId, s.ServiceToken)

	if s.DeviceId != "" && s.SerialNumber != "" {
		cookie = fmt.Sprintf("%s;deviceId=%s;sn=%s", cookie, s.UserId, s.ServiceToken)
	}
	return cookie
}
