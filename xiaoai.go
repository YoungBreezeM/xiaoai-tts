package xiaoaitts

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func LoginByAccount(m *MiAccount) *Session {
Login:
	sign := GetLoginSign()
	authInfo := ServiceAuth(sign.Sign, m)
	//
	if authInfo.Code != 0 {
		fmt.Println("errr")
	}
	//
	session := &Session{}
	token := LoginMiAi(authInfo)
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

type XiaoAiFunc interface {
	GetDevice() []DeviceInfo
	UseDevice(index int16)
	Say(text string)
	SetVolume(volume int8)
	GetVolume() string
	Play()
	Pause()
	Prev()
	Next()
	TogglePlayState()
	GetStatus() *Info
	PlayUrl(url string)
}

type XiaoAi struct {
	Session *Session
}

func NewXiaoAi(m *MiAccount) XiaoAiFunc {
	x := &XiaoAi{}
	x.Session = LoginByAccount(m)
	//Default switch first device
	msg := GetDevice(x.Session)
	//
	if msg.Data != nil {
		device := msg.Data[0]
		x.Session.DeviceId = device.DeviceID
		x.Session.SerialNumber = device.SerialNumber
	}
	//
	return x
}

func (x *XiaoAi) GetDevice() []DeviceInfo {
	return GetDevice(x.Session).Data
}

func (x *XiaoAi) UseDevice(index int16) {
	device := GetDevice(x.Session).Data[index]
	x.Session.DeviceId = device.DeviceID
	x.Session.SerialNumber = device.SerialNumber
}

func (x *XiaoAi) Say(text string) {
	msg, _ := json.Marshal(Message{
		Text:  text,
		Save:  0,
		Media: "app_ios",
	})
	Ubus(&Ticket{
		Cookie:   x.Session.GetCookie(),
		DeviceId: x.Session.DeviceId,
	}, &UbusParam{
		Method:  "text_to_speech",
		Message: string(msg),
		Path:    "mibrain",
	})
}

func (x *XiaoAi) SetVolume(volume int8) {
	msg, _ := json.Marshal(Message{
		Volume: volume,
	})
	Ubus(&Ticket{
		Cookie:   x.Session.GetCookie(),
		DeviceId: x.Session.DeviceId,
	}, &UbusParam{
		Method:  "player_set_volume",
		Message: string(msg),
		Path:    "mediaplayer",
	})
}

func (x *XiaoAi) GetVolume() string {
	msg, _ := json.Marshal(Message{})
	res := Ubus(&Ticket{
		Cookie:   x.Session.GetCookie(),
		DeviceId: x.Session.DeviceId,
	}, &UbusParam{
		Method:  "player_get_play_status",
		Message: string(msg),
		Path:    "mediaplayer",
	})
	return ParseVolume(string(res))
}

func (x *XiaoAi) Play() {
	msg, _ := json.Marshal(Message{
		Action: "play",
	})
	//
	Ubus(&Ticket{
		Cookie:   x.Session.GetCookie(),
		DeviceId: x.Session.DeviceId,
	}, &UbusParam{
		Method:  "player_play_operation",
		Message: string(msg),
		Path:    "mediaplayer",
	})

}

func (x *XiaoAi) Pause() {
	msg, _ := json.Marshal(Message{
		Action: "pause",
	})
	//
	Ubus(&Ticket{
		Cookie:   x.Session.GetCookie(),
		DeviceId: x.Session.DeviceId,
	}, &UbusParam{
		Method:  "player_play_operation",
		Message: string(msg),
		Path:    "mediaplayer",
	})

}

func (x *XiaoAi) Prev() {
	msg, _ := json.Marshal(Message{
		Action: "prev",
	})
	//
	Ubus(&Ticket{
		Cookie:   x.Session.GetCookie(),
		DeviceId: x.Session.DeviceId,
	}, &UbusParam{
		Method:  "player_play_operation",
		Message: string(msg),
		Path:    "mediaplayer",
	})
}

func (x *XiaoAi) Next() {
	msg, _ := json.Marshal(Message{
		Action: "next",
	})
	//
	Ubus(&Ticket{
		Cookie:   x.Session.GetCookie(),
		DeviceId: x.Session.DeviceId,
	}, &UbusParam{
		Method:  "player_play_operation",
		Message: string(msg),
		Path:    "mediaplayer",
	})
}

func (x *XiaoAi) TogglePlayState() {
	msg, _ := json.Marshal(Message{
		Action: "toggle",
	})
	//
	Ubus(&Ticket{
		Cookie:   x.Session.GetCookie(),
		DeviceId: x.Session.DeviceId,
	}, &UbusParam{
		Method:  "player_play_operation",
		Message: string(msg),
		Path:    "mediaplayer",
	})
}

func (x *XiaoAi) GetStatus() *Info {
	msg, _ := json.Marshal(Message{})
	//
	res := Ubus(&Ticket{
		Cookie:   x.Session.GetCookie(),
		DeviceId: x.Session.DeviceId,
	}, &UbusParam{
		Method:  "player_get_play_status",
		Message: string(msg),
		Path:    "mediaplayer",
	})
	fmt.Println(string(res))
	c := regexp.MustCompile("\"info\":\"(.*)\"}}")
	s := c.FindStringSubmatch(string(res))
	s2 := s[len(s)-1]
	s3 := strings.Replace(s2, "\\", "", -1)
	m := &Info{}
	err := json.Unmarshal([]byte(s3), m)
	if err != nil {
		log.Print(err)
	}
	return m
}

func (x *XiaoAi) PlayUrl(url string) {
	msg, _ := json.Marshal(Message{
		Url:   url,
		Media: "app_ios",
		Type:  1,
	})
	//
	Ubus(&Ticket{
		Cookie:   x.Session.GetCookie(),
		DeviceId: x.Session.DeviceId,
	}, &UbusParam{
		Method:  "player_play_url",
		Message: string(msg),
		Path:    "mediaplayer",
	})

}
