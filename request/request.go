package request

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	api_url                   string = "https://api.switch-bot.com"
	api_version               string = "v1.0"
	path_device               string = "devices"
	header_auth               string = "Authorization"
	header_content_type       string = "Content-Type"
	header_value_content_type string = "application/json"
)

type (
	Swichgot struct {
		token string
	}
	Device struct {
		id                 string
		name               string
		deviceType         string
		enableCloudService *bool
		hubID              string
	}
)

func NewClient(token string) *Swichgot {
	return &Swichgot{token: token}
}

func (s *Swichgot) List(show bool) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", api_url, api_version, path_device), nil)
	if err != nil {
		log.Fatalf("Request create error. %s", err)
	}
	req.Header.Set(header_auth, s.token)
	req.Header.Set(header_content_type, header_value_content_type)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Request failed. %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Response not OK. have %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Response read falied %s", err)
	}
	var obj interface{}
	err = json.Unmarshal(body, &obj)
	if err != nil {
		log.Fatalf("Body decode failed %s", err)
	}
	reqlist := obj.(map[string]interface{})["body"].(map[string]interface{})["deviceList"].([]interface{})
	list := make([]Device, len(reqlist))
	for i, v := range reqlist {
		list[i].id = v.(map[string]interface{})["deviceId"].(string)
		list[i].name = v.(map[string]interface{})["deviceName"].(string)
		list[i].deviceType = v.(map[string]interface{})["deviceType"].(string)
		if v.(map[string]interface{})["enableCloudService"] != nil {
			en := v.(map[string]interface{})["enableCloudService"]
			enable := en.(bool)
			list[i].enableCloudService = &enable
		}
		list[i].hubID = v.(map[string]interface{})["hubDeviceId"].(string)
	}
	if show {
		s := "[\n"
		for i, v := range list {
			s += fmt.Sprintf("\t{\n\t\t\"id\": \"%s\",\n\t\t\"name\": \"%s\",\n\t\t\"deviceType\": \"%s\",\n", v.id, v.name, v.deviceType)
			if v.enableCloudService != nil {
				s += fmt.Sprintf("\t\t\"enableCloudService\": %t,\n", *v.enableCloudService)
			}
			s += fmt.Sprintf("\t\t\"hubID\": \"%s\"\n\t}", v.hubID)
			if i != len(list)-1 {
				s += ","
			}
			s += "\n"
		}
		s += "]"
		fmt.Println(s)
	}
}
