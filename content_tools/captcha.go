package content_tools

import (
	"crypto/tls"
	"encoding/json"
	"github.com/kmahyyg/pb-go/config"
	"github.com/valyala/fasthttp"
	"net/url"
	"time"
)

const recaptchaServUrl = "https://www.google.com/recaptcha/api/siteverify"

type ReCaptchaResponse struct {
	Success     bool      `json:"success"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

func VerifyRecaptcha(recaptchaResponse string, remoteIp string) (bool, error) {
	var err error
	httpc := fasthttp.Client{
		NoDefaultUserAgentHeader: true,
		TLSConfig:                &tls.Config{InsecureSkipVerify:false},
		ReadTimeout:              5 * time.Second,
		WriteTimeout:             5 * time.Second,
	}
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(recaptchaServUrl)
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBodyString(url.Values{
		"secret":   {config.ServConf.Recaptcha.Secret_key},
		"response": {recaptchaResponse},
		"remoteip": {remoteIp},
	}.Encode())
	resp := fasthttp.AcquireResponse()
	if err = httpc.Do(req, resp); err != nil {
		return false, err
	}
	recaptResp := ReCaptchaResponse{}
	if resp.StatusCode() == fasthttp.StatusOK {
		err = json.Unmarshal(resp.Body(), &recaptResp)
		res := recaptResp.Success
		return res, err
	}
	return false, err
}
