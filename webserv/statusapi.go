package webserv

import (
	"encoding/json"
	"github.com/pb-go/pb-go/config"
	"log"
)

// ServStatusData : Json Serialize data, please note, AllowBase64Encode Require FileCheckEnabled set to true.
type ServStatusData struct {
	RunHealth string `json:"status"`
	CaptchaEnabled bool `json:"captcha_enabled"`
	MaxExpireTime int `json:"max_expire"`
	FileCheckEnabled bool `json:"abuse_detection"`
	AllowBase64Encode bool `json:"base64_detection"`
}


// retStatusJson : Return Server Public Configuration for Client
func retStatusJson() []byte{
	servstat := ServStatusData{
		RunHealth:		   "Running",
		CaptchaEnabled:    config.ServConf.Recaptcha.Enable,
		MaxExpireTime:     config.ServConf.Content.ExpireHrs,
		FileCheckEnabled:  config.ServConf.Content.DetectAbuse,
		AllowBase64Encode: config.ServConf.Content.AllowBase64Encode,
	}
	renderedstat, err := json.Marshal(servstat)
	if err != nil {
		log.Fatalln("Failed to generate Server Status Page.")
	}
	return renderedstat
}
