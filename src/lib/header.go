package lib

import (
	"net/http"
	"strings"

	"github.com/qfyang-cn/xiaoai-tts/src/constant"
)

func SetXiaoAiRequestHeaders(request *http.Request) {
	contentType := "application/json"
	if request.Method == http.MethodPost {
		contentType = "application/x-www-form-urlencoded"
	}
	//
	userAgent := constant.APP_UA
	if strings.Contains(request.URL.Host, "mina.mi.com") {
		userAgent = constant.MINA_UA
	}
	//
	request.Header.Add("Content-Type", contentType)
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("User-Agent", userAgent)
	request.Header.Add("Accept", "*/*")
}
