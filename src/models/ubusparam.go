package models

type Message struct {
	Text   string `url:"text" json:"text"`
	Save   int8   `url:"save" json:"save"`
	Media  string `url:"media" json:"media"`
	Volume int8   `url:"volume" json:"volume"`
	Action string `url:"action" json:"action"`
	Url    string `url:"url" json:"url"`
	Type   int8   `url:"type" json:"type"`
}

type UbusParam struct {
	Method    string `url:"method" `
	Message   string `url:"message" `
	Path      string `url:"path" `
	RequestId string `url:"requestId" `
	DeviceId  string `url:"deviceId" `
}
