package clipkg

import (
	"github.com/pb-go/pb-go/config"
	"github.com/valyala/fasthttp"
)

func MakeRequest() *fasthttp.Request {
	request := fasthttp.AcquireRequest()
	request.Header.SetUserAgent("pb-cli/" + config.CurrentVer)
	return request
}
