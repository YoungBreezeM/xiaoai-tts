package req

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/qfyang-cn/xiaoai-tts/src/constant"
	"github.com/qfyang-cn/xiaoai-tts/src/models"
	"github.com/qfyang-cn/xiaoai-tts/src/utils"

	"github.com/google/go-querystring/query"
)

func Ubus(t *models.Ticket, p *models.UbusParam) []byte {
	p.DeviceId = t.DeviceId
	p.RequestId = fmt.Sprintf("app_ios_%s", utils.GetRandomString(30))
	//
	v, _ := query.Values(p)
	//
	req, err := http.NewRequest(http.MethodPost, constant.USBS, nil)
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
	body := utils.ParseResponse(resp)
	return body
}
