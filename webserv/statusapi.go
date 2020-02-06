package webserv

import (
	"encoding/json"
	"github.com/pb-go/pb-go/config"
	"log"
)

// ServStatusData : Json Serialize data about server config
type ServStatusData struct {
	RunHealth        int  `json:"status"`
	CaptchaEnabled   bool `json:"captcha_enabled"`
	MaxExpireTime    int  `json:"max_expire"`
	FileCheckEnabled bool `json:"abuse_detection"`
}

// retStatusJson : Return Server Public Configuration for Client
func retStatusJSON() []byte {
	servstat := ServStatusData{
		RunHealth:        0,
		CaptchaEnabled:   config.ServConf.Recaptcha.Enable,
		MaxExpireTime:    config.ServConf.Content.ExpireHrs,
		FileCheckEnabled: config.ServConf.Content.DetectAbuse,
	}
	renderedstat, err := json.Marshal(servstat)
	if err != nil {
		log.Fatalln("Failed to generate Server Status Page.")
	}
	return renderedstat
}
