package models

import (
	"encoding/json"
	"fmt"
)

type ServiceParam struct {
	CheckSafePhone   bool
	CheckSafeAddress bool
	Lsrp_score       float32
}

type LoginSign struct {
	ServiceParam ServiceParam
	*Sign
	Code           int16
	Description    string
	SecurityStatus int16
	Sid            string
	Result         string
	CaptchaUrl     string
	Callback       string
	Location       string
	Pwd            int16
	Child          int16
	Desc           string
}

type Sign struct {
	Sign string `url:"_sign" json:"_sign"`
	Qs   string `url:"qs" json:"qs"`
}

func (login *LoginSign) String() string {
	v, err := json.MarshalIndent(login, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	return string(v)
}
