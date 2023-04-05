package req

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/qfyang-cn/xiaoai-tts/constant"
	"github.com/qfyang-cn/xiaoai-tts/models"
	"github.com/qfyang-cn/xiaoai-tts/utils"
)

func GetDevice(s *models.Session) *models.Msg {
	req, _ := http.NewRequest(http.MethodGet, constant.DEVICE_LIST, nil)
	req.Header.Add("Cookie", s.GetCookie())
	req.URL.RawQuery = fmt.Sprintf("master=1&requestId=%s", utils.GetRandomString(30))
	//
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()
	//
	body, _ := ioutil.ReadAll(resp.Body)
	m := &models.Msg{}
	json.Unmarshal(body, m)
	//
	return m
}

func GetSongInfo(cookie string, songId string) []byte {
	req, _ := http.NewRequest(http.MethodGet, constant.SONG_INFO, nil)
	req.Header.Add("Cookie", cookie)
	req.URL.RawQuery = fmt.Sprintf("songId=%s", songId)
	//
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()
	//
	body, _ := ioutil.ReadAll(resp.Body)

	//
	return body
}
