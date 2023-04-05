package service

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"xiaoai-tts/src/models"
	"xiaoai-tts/src/req"
	"xiaoai-tts/src/utils"
)

type XiaoAiFunc interface {
	GetDevice() []models.DeviceInfo
	UseDevice(index int16)
	Say(text string)
	SetVolume(volume int8)
	GetVolume() string
	Play()
	Pause()
	Prev()
	Next()
	TogglePlayState()
	GetStatus() *models.Info
	PlayUrl(url string)
}

type XiaoAi struct {
	Session *models.Session
}

func NewXiaoAi(m *models.MiAccount) XiaoAiFunc {
	x := &XiaoAi{}
	x.Session = Login(m)
	//Default switch first device
	msg := req.GetDevice(x.Session)
	//
	if msg.Data != nil {
		device := msg.Data[0]
		x.Session.DeviceId = device.DeviceID
		x.Session.SerialNumber = device.SerialNumber
	}
	//
	return x
}

func (x *XiaoAi) GetDevice() []models.DeviceInfo {
	return req.GetDevice(x.Session).Data
}

func (x *XiaoAi) UseDevice(index int16) {
	device := req.GetDevice(x.Session).Data[index]
	x.Session.DeviceId = device.DeviceID
	x.Session.SerialNumber = device.SerialNumber
}

func (x *XiaoAi) Say(text string) {
	msg, _ := json.Marshal(models.Message{
		Text:  text,
		Save:  0,
		Media: "app_ios",
	})
	req.Ubus(&models.Ticket{
		Cookie:   x.Session.GetCookie(),
		DeviceId: x.Session.DeviceId,
	}, &models.UbusParam{
		Method:  "text_to_speech",
		Message: string(msg),
		Path:    "mibrain",
	})
}

func (x *XiaoAi) SetVolume(volume int8) {
	msg, _ := json.Marshal(models.Message{
		Volume: volume,
	})
	req.Ubus(&models.Ticket{
		Cookie:   x.Session.GetCookie(),
		DeviceId: x.Session.DeviceId,
	}, &models.UbusParam{
		Method:  "player_set_volume",
		Message: string(msg),
		Path:    "mediaplayer",
	})
}

func (x *XiaoAi) GetVolume() string {
	msg, _ := json.Marshal(models.Message{})
	res := req.Ubus(&models.Ticket{
		Cookie:   x.Session.GetCookie(),
		DeviceId: x.Session.DeviceId,
	}, &models.UbusParam{
		Method:  "player_get_play_status",
		Message: string(msg),
		Path:    "mediaplayer",
	})
	return utils.ParseVolume(string(res))
}

func (x *XiaoAi) Play() {
	msg, _ := json.Marshal(models.Message{
		Action: "play",
	})
	//
	req.Ubus(&models.Ticket{
		Cookie:   x.Session.GetCookie(),
		DeviceId: x.Session.DeviceId,
	}, &models.UbusParam{
		Method:  "player_play_operation",
		Message: string(msg),
		Path:    "mediaplayer",
	})

}

func (x *XiaoAi) Pause() {
	msg, _ := json.Marshal(models.Message{
		Action: "pause",
	})
	//
	req.Ubus(&models.Ticket{
		Cookie:   x.Session.GetCookie(),
		DeviceId: x.Session.DeviceId,
	}, &models.UbusParam{
		Method:  "player_play_operation",
		Message: string(msg),
		Path:    "mediaplayer",
	})

}

func (x *XiaoAi) Prev() {
	msg, _ := json.Marshal(models.Message{
		Action: "prev",
	})
	//
	req.Ubus(&models.Ticket{
		Cookie:   x.Session.GetCookie(),
		DeviceId: x.Session.DeviceId,
	}, &models.UbusParam{
		Method:  "player_play_operation",
		Message: string(msg),
		Path:    "mediaplayer",
	})
}

func (x *XiaoAi) Next() {
	msg, _ := json.Marshal(models.Message{
		Action: "next",
	})
	//
	req.Ubus(&models.Ticket{
		Cookie:   x.Session.GetCookie(),
		DeviceId: x.Session.DeviceId,
	}, &models.UbusParam{
		Method:  "player_play_operation",
		Message: string(msg),
		Path:    "mediaplayer",
	})
}

func (x *XiaoAi) TogglePlayState() {
	msg, _ := json.Marshal(models.Message{
		Action: "toggle",
	})
	//
	req.Ubus(&models.Ticket{
		Cookie:   x.Session.GetCookie(),
		DeviceId: x.Session.DeviceId,
	}, &models.UbusParam{
		Method:  "player_play_operation",
		Message: string(msg),
		Path:    "mediaplayer",
	})
}

func (x *XiaoAi) GetStatus() *models.Info {
	msg, _ := json.Marshal(models.Message{})
	//
	res := req.Ubus(&models.Ticket{
		Cookie:   x.Session.GetCookie(),
		DeviceId: x.Session.DeviceId,
	}, &models.UbusParam{
		Method:  "player_get_play_status",
		Message: string(msg),
		Path:    "mediaplayer",
	})
	fmt.Println(string(res))
	c := regexp.MustCompile("\"info\":\"(.*)\"}}")
	s := c.FindStringSubmatch(string(res))
	s2 := s[len(s)-1]
	s3 := strings.Replace(s2, "\\", "", -1)
	m := &models.Info{}
	json.Unmarshal([]byte(s3), m)
	return m
}

func (x *XiaoAi) PlayUrl(url string) {
	msg, _ := json.Marshal(models.Message{
		Url:   url,
		Media: "app_ios",
		Type:  1,
	})
	//
	req.Ubus(&models.Ticket{
		Cookie:   x.Session.GetCookie(),
		DeviceId: x.Session.DeviceId,
	}, &models.UbusParam{
		Method:  "player_play_url",
		Message: string(msg),
		Path:    "mediaplayer",
	})

}
