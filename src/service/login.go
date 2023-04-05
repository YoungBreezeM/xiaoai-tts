package service

import (
	"fmt"
	"strconv"

	"github.com/qfyang-cn/xiaoai-tts/src/models"
	"github.com/qfyang-cn/xiaoai-tts/src/req"
)

func LoginByAccount(m *models.MiAccount) *models.Session {
Login:
	sign := req.GetLoginSign()
	authInfo := req.ServiceAuth(sign.Sign, m)
	//
	if authInfo.Code != 0 {
		fmt.Println("errr")
	}
	//
	session := &models.Session{}
	token := req.LoginMiAi(authInfo)
	//
	if len(token) > 0 {
		session.ServiceToken = token
		session.UserId = strconv.Itoa(authInfo.UserID)
	} else {
		goto Login
	}
	//
	return session
}

func Login(m *models.MiAccount) *models.Session {
	s := &models.Session{}
	s = LoginByAccount(m)
	return s
}
