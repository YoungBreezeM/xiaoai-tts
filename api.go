package xiaoaitts

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	USBS           = "https://api.mina.mi.com/remote/ubus"
	SERVICE_AUTH   = "https://account.xiaomi.com/pass/serviceLoginAuth2"
	SERVICE_LOGIN  = "https://account.xiaomi.com/pass/serviceLogin"
	PLAYLIST       = "https://api2.mina.mi.com/music/playlist/v2/lists"
	PLAYLIST_SONGS = "https://api2.mina.mi.com/music/playlist/v2/songs"
	DEVICE_LIST    = "https://api.mina.mi.com/admin/v2/device_list"
	SONG_INFO      = "https://api2.mina.mi.com/music/song_info"
	APP_DEVICE_ID  = "3C861A5820190429"
	SDK_VER        = "3.4.1"
	APP_UA         = "APP/com.xiaomi.mico APPV/2.1.17 iosPassportSDK/3.4.1 iOS/13.3.1"
	MINA_UA        = "MISoundBox/2.1.17 (com.xiaomi.mico; build:2.1.55; iOS 13.3.1) Alamofire/4.8.2 MICO/iOSApp/appStore/2.1.17"
	SID            = "micoapi"
	JSON           = "true"
)

func SetXiaoAiRequestHeaders(request *http.Request) {
	contentType := "application/json"
	if request.Method == http.MethodPost {
		contentType = "application/x-www-form-urlencoded"
	}
	//
	userAgent := APP_UA
	if strings.Contains(request.URL.Host, "mina.mi.com") {
		userAgent = MINA_UA
	}
	//
	request.Header.Add("Content-Type", contentType)
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("User-Agent", userAgent)
	request.Header.Add("Accept", "*/*")
}

func GetLoginSign() *LoginSign {
	req, err := http.NewRequest(http.MethodGet, SERVICE_LOGIN, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	SetXiaoAiRequestHeaders(req)
	//
	req.URL.RawQuery = fmt.Sprintf("sid=%s&_json=%s", SID, JSON)
	//
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()
	//
	sign := &LoginSign{}
	d := ParseResponse(resp)
	err = json.Unmarshal(d, sign)
	if err != nil {
		log.Print(err)
	}
	//
	return sign
}

func ServiceAuth(sign *Sign, m *MiAccount) *Auth {
	hash := md5.Sum([]byte(m.Pwd))
	hashStr := fmt.Sprintf("%x", hash)
	authData := &AuthData{
		User:     m.User,
		Hash:     strings.ToUpper(hashStr),
		Callback: "https://api.mina.mi.com/sts",
		CommonParam: &CommonParam{
			Sid:  SID,
			Json: JSON,
		},
		Sign: sign,
	}
	//
	v, _ := query.Values(authData)
	payload := bytes.NewReader([]byte(v.Encode()))
	//
	req, err := http.NewRequest(http.MethodPost, SERVICE_AUTH, payload)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	//
	req.Header.Add("Cookie", fmt.Sprintf("deviceId=%s;sdkVersion=%s", APP_DEVICE_ID, SDK_VER))
	req.Header.Add("Content-Length", strconv.Itoa(len(v.Encode())))
	//
	SetXiaoAiRequestHeaders(req)
	//
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()
	//
	body := ParseResponse(resp)
	authInfo := &Auth{}
	err = json.Unmarshal(body, authInfo)
	if err != nil {
		log.Print(err)
	}
	return authInfo
}

func LoginMiAi(authInfo *Auth) string {
	signStr := fmt.Sprintf("nonce=%s&%s", strconv.Itoa(authInfo.Nonce), authInfo.Ssecurity)
	clientSign := Sha1Base64(signStr)
	authInfo.Location = fmt.Sprintf("%s&clientSign=%s", authInfo.Location, clientSign)
	//
	req, err := http.NewRequest(http.MethodGet, authInfo.Location, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	//
	SetXiaoAiRequestHeaders(req)
	//
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()
	token := ""
	if resp.StatusCode == 200 {
		token = ParseToekn(resp.Header.Get("Set-Cookie"))
	}
	//
	return token
}

func GetDevice(s *Session) *Msg {
	req, _ := http.NewRequest(http.MethodGet, DEVICE_LIST, nil)
	req.Header.Add("Cookie", s.GetCookie())
	req.URL.RawQuery = fmt.Sprintf("master=1&requestId=%s", GetRandomString(30))
	//
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()
	//
	body, _ := io.ReadAll(resp.Body)
	m := &Msg{}
	err = json.Unmarshal(body, m)
	if err != nil {
		log.Print(err)
	}
	//
	return m
}

func GetSongInfo(cookie string, songId string) []byte {
	req, _ := http.NewRequest(http.MethodGet, SONG_INFO, nil)
	req.Header.Add("Cookie", cookie)
	req.URL.RawQuery = fmt.Sprintf("songId=%s", songId)
	//
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()
	//
	body, _ := io.ReadAll(resp.Body)

	//
	return body
}

func Ubus(t *Ticket, p *UbusParam) []byte {
	p.DeviceId = t.DeviceId
	p.RequestId = fmt.Sprintf("app_ios_%s", GetRandomString(30))
	//
	v, _ := query.Values(p)
	//
	req, err := http.NewRequest(http.MethodPost, USBS, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	//
	req.URL.RawQuery = v.Encode()
	req.Header.Add("Cookie", t.Cookie)
	//
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()
	//
	body := ParseResponse(resp)
	return body
}
