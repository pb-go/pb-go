package contenttools

import (
	"crypto/tls"
	"encoding/json"
	"github.com/pb-go/pb-go/config"
	"github.com/valyala/fasthttp"
	"net/url"
	"time"
)

const recaptchaServURL = "https://www.google.com/recaptcha/api/siteverify"

// ReCaptchaResponse : defined by google, returns as json
type ReCaptchaResponse struct {
	Success     bool      `json:"success"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

// VerifyRecaptchaResp : POST user IP and response token to do SSV
func VerifyRecaptchaResp(recaptchaResponse string, remoteIP string) (bool, error) {
	var err error
	httpc := fasthttp.Client{
		NoDefaultUserAgentHeader: true,
		TLSConfig:                &tls.Config{InsecureSkipVerify: false},
		ReadTimeout:              5 * time.Second,
		WriteTimeout:             5 * time.Second,
	}
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(recaptchaServURL)
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBodyString(url.Values{
		"secret":   {config.ServConf.Recaptcha.SecretKey},
		"response": {recaptchaResponse},
		"remoteip": {remoteIP},
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
