package req

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/qfyang-cn/xiaoai-tts/src/constant"
	"github.com/qfyang-cn/xiaoai-tts/src/lib"
	"github.com/qfyang-cn/xiaoai-tts/src/models"
	"github.com/qfyang-cn/xiaoai-tts/src/utils"
)

var commonParam = &models.CommonParam{
	Sid:  "micoapi",
	Json: "true",
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////
func GetLoginSign() *models.LoginSign {
	req, err := http.NewRequest(http.MethodGet, constant.SERVICE_LOGIN, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	lib.SetXiaoAiRequestHeaders(req)
	//
	req.URL.RawQuery = fmt.Sprintf("sid=%s&_json=%s", commonParam.Sid, commonParam.Json)
	//
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()
	//
	sign := &models.LoginSign{}
	d := utils.ParseResponse(resp)
	json.Unmarshal(d, sign)
	//
	return sign
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////
func ServiceAuth(sign *models.Sign, m *models.MiAccount) *models.Auth {
	hash := md5.Sum([]byte(m.Pwd))
	hashStr := fmt.Sprintf("%x", hash)
	authData := &models.AuthData{
		User:        m.User,
		Hash:        strings.ToUpper(hashStr),
		Callback:    "https://api.mina.mi.com/sts",
		CommonParam: commonParam,
		Sign:        sign,
	}
	//
	v, _ := query.Values(authData)
	payload := bytes.NewReader([]byte(v.Encode()))
	//
	req, err := http.NewRequest(http.MethodPost, constant.SERVICE_AUTH, payload)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	//
	req.Header.Add("Cookie", fmt.Sprintf("deviceId=%s;sdkVersion=%s", constant.APP_DEVICE_ID, constant.SDK_VER))
	req.Header.Add("Content-Length", strconv.Itoa(len(v.Encode())))
	//
	lib.SetXiaoAiRequestHeaders(req)
	//
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()
	//
	body := utils.ParseResponse(resp)
	authInfo := &models.Auth{}
	json.Unmarshal(body, authInfo)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	return authInfo
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////
func LoginMiAi(authInfo *models.Auth) string {
	signStr := fmt.Sprintf("nonce=%s&%s", strconv.Itoa(authInfo.Nonce), authInfo.Ssecurity)
	clientSign := utils.Sha1Base64(signStr)
	authInfo.Location = fmt.Sprintf("%s&clientSign=%s", authInfo.Location, clientSign)
	//
	req, err := http.NewRequest(http.MethodGet, authInfo.Location, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	//
	lib.SetXiaoAiRequestHeaders(req)
	//
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()
	token := ""
	if resp.StatusCode == 200 {
		token = utils.ParseToekn(resp.Header.Get("Set-Cookie"))
	}
	//
	return token
}
