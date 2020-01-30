package content_tools

import (
	"encoding/json"
	"github.com/kmahyyg/pb-go/config"
	"net/http"
	"net/url"
	"time"
)

const recaptchaUrl = "https://www.google.com/recaptcha/api/siteverify"

type ReCaptchaResponse struct {
	Success     bool      `json:"success"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

func verifyRecaptcha(config config.Recaptcha, recaptchaResponse string, remoteIp string) (result bool, err error) {
	response, err := http.PostForm(recaptchaUrl, url.Values{
		"secret":   {config.Secret_key},
		"response": {recaptchaResponse},
		"remoteip": {remoteIp},
	})
	if err == nil {
		defer response.Body.Close()
		reCaptchaResponse := ReCaptchaResponse{}
		err = json.NewDecoder(response.Body).Decode(reCaptchaResponse)
		result = reCaptchaResponse.Success
		return
	}
	return
}
