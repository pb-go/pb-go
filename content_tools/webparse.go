package content_tools

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func UserUploadParse(c *fasthttp.RequestCtx) {
	// todo: remove to use fasthttp as replace
	panic("todo: not implemented")
}

func ShowSnip(c *fasthttp.RequestCtx) {
	// todo: remove to use fasthttp as replace
	tmpvar := c.UserValue("shortId")
	switch tmpvar {
	case nil:
		fallthrough
	case "index.html":
		fasthttp.ServeFile(c, "./static/index.html")
	case "submit.html":
		fasthttp.ServeFile(c, "./static/submit.html")
	case "favicon.ico":
		fasthttp.ServeFile(c, "./static/favicon.ico")
	default:
		fmt.Println(tmpvar)
	}
}


func DeleteSnip(c *fasthttp.RequestCtx) {
	// todo: remove to use fasthttp as replace
	panic("todo: not implemented")
}

func VerifyCAPT(c *fasthttp.RequestCtx) {
	// todo: remove to use fasthttp as replace
	panic("todo: not implemented")
}
