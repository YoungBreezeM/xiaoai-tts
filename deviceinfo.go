package xiaoaitts

import "encoding/json"

type DeviceInfo struct {
	DeviceID        string           `json:"deviceID"`
	SerialNumber    string           `json:"serialNumber"`
	Name            string           `json:"name"`
	Alias           string           `json:"alias"`
	Current         bool             `json:"current"`
	Presence        string           `json:"presence"`
	Address         string           `json:"address"`
	MiotDID         string           `json:"miotDID"`
	Hardware        string           `json:"hardware"`
	ROMVersion      string           `json:"romVersion"`
	Capabilities    map[string]int64 `json:"capabilities"`
	RemoteCtrlType  string           `json:"remoteCtrlType"`
	DeviceSNProfile string           `json:"deviceSNProfile"`
	DeviceProfile   string           `json:"deviceProfile"`
	BrokerEndpoint  string           `json:"brokerEndpoint"`
	BrokerIndex     int64            `json:"brokerIndex"`
	MAC             string           `json:"mac"`
	SSID            string           `json:"ssid"`
}

func (d *DeviceInfo) String() string {
	s, _ := json.MarshalIndent(d, "", " ")
	return string(s)
}
