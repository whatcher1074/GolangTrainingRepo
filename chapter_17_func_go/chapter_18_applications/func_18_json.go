package main

import (
	"encoding/json"
	"fmt"
)

type SMBCable struct {
	DevicePresent     bool   `json:"devicePresent"`
	DeviceID          string `json:"deviceId"`
	DeviceType        string `json:"deviceType"`
	CmMacAddress      string `json:"cmMacAddress"`
	MtaMacAddress     string `json:"mtaMacAddress"`
	WifiMacAddress    string `json:"wifiMacAddress"`
	EcmMacAddress     string `json:"ecmMacAddress"`
	Make              string `json:"make"`
	Model             string `json:"model"`
	AssignmentStatus  string `json:"assignmentStatus"`
	NewDeviceInBiller bool   `json:"newDeviceInBiller"`
	NewDevice         bool   `json:"newDevice"`
	CustomerOwned     bool   `json:"customerOwned"`
	ScpRouter         bool   `json:"scpRouter"`
	HasServices       bool   `json:"hasServices"`
	RecentDisconnect  bool   `json:"recentDisconnect"`
	DeviceRole        string `json:"deviceRole"`
	DvrCapability     bool   `json:"dvrCapability"`
	HdCapability      bool   `json:"hdCapability"`
	PendingServices   []struct {
		ServiceName      string `json:"serviceName"`
		ServiceStatus    string `json:"serviceStatus"`
		LineOfBusiness   string `json:"lineOfBusiness"`
		Description      string `json:"description"`
		WifiServiceID    string `json:"wifiServiceId"`
		StreamingService bool   `json:"streamingService"`
		MaxDeviceCount   int    `json:"maxDeviceCount"`
		MaxPodCount      int    `json:"maxPodCount"`
	} `json:"pendingServices"`

	ActiveServices []struct {
		ServiceName      string `json:"serviceName"`
		ServiceStatus    string `json:"serviceStatus"`
		LineOfBusiness   string `json:"lineOfBusiness"`
		Description      string `json:"description"`
		WifiServiceID    string `json:"wifiServiceId"`
		StreamingService bool   `json:"streamingService"`
		MaxDeviceCount   int    `json:"maxDeviceCount"`
		MaxPodCount      int    `json:"maxPodCount"`
	} `json:"activeServices"`
	AwgDevice bool `json:"awgDevice"`
}

func main() {

	activationData := `[{
		"devicePresent": true,
		"deviceId": "E448C71912B0",
		"deviceType": "Router",
		"cmMacAddress": "",
		"mtaMacAddress": "",
		"wifiMacAddress": "E448C71912B0",
		"ecmMacAddress": "",
		"make": "ASKEY",
		"model": "RAC2V1K",
		"assignmentStatus": "ACTIVE",
		"newDeviceInBiller": false,
		"newDevice": false,
		"customerOwned": false,
		"scpRouter": false,
		"hasServices": true,
		"recentDisconnect": false,
		"deviceRole": "NONE",
		"dvrCapability": false,
		"hdCapability": false,
		"pendingServices": [
		  {
			"serviceName": "1101521837794_8087300230002977_000001",
			"serviceStatus": "PENDING",
			"lineOfBusiness": "WIFI",
			"description": "Spectrum HTSP-AP",
			"wifiServiceId": "000001",
			"streamingService": false,
			"maxDeviceCount": 0,
			"maxPodCount": 0
		  }
		],
		"activeServices": [
		  {
			"serviceName": "1101521837794_8087300230002977_000010",
			"serviceStatus": "ACTIVE",
			"lineOfBusiness": "WIFI",
			"description": "Spectrum WiFi",
			"wifiServiceId": "000010",
			"streamingService": false,
			"maxDeviceCount": 0,
			"maxPodCount": 0
		  }
		],
		"awgDevice": false
	  }]`

	bs := []byte(activationData)
	fmt.Printf("%T\n", bs)

	var cableRequest []SMBCable

	err := json.Unmarshal(bs, &cableRequest)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cableRequest)
	cableDeviceID := cableRequest[0].DeviceID
	fmt.Println(cableRequest[0].DevicePresent)
	fmt.Println(cableDeviceID)

}
