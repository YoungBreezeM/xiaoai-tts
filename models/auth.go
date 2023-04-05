package models

import "encoding/json"

type Auth struct {
	Qs             string `json:"qs"`
	Ssecurity      string `json:"ssecurity"`
	Code           int    `json:"code"`
	PassToken      string `json:"passToken"`
	Description    string `json:"description"`
	SecurityStatus int    `json:"securityStatus"`
	Nonce          int    `json:"nonce"`
	UserID         int    `json:"userId"`
	CUserID        string `json:"cUserId"`
	Result         string `json:"result"`
	Psecurity      string `json:"psecurity"`
	CAPTCHAURL     string `json:"captchaUrl"`
	Location       string `json:"location"`
	Pwd            int    `json:"pwd"`
	Child          int    `json:"child"`
	Desc           string `json:"desc"`
}

func UnmarshalAuth(data []byte) (Auth, error) {
	var r Auth
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Auth) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Auth) String() string {
	v, _ := json.MarshalIndent(r, "", " ")
	return string(v)
}

////////////////////////////////////////////////////
type AuthData struct {
	User     string `url:"user"`
	Hash     string `url:"hash"`
	Callback string `url:"callback"`
	*CommonParam
	*Sign
}
