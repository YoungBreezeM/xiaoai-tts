package models

import "encoding/json"

type Msg struct {
	Code    int64        `json:"code"`
	Message string       `json:"message"`
	Data    []DeviceInfo `json:"data"`
}

func (m *Msg) String() string {
	s, _ := json.MarshalIndent(m, "", " ")
	return string(s)
}
