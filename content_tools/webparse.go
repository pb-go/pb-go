package content_tools

import (
	"fmt"
	"log"
	"github.com/valyala/fasthttp"
)

func UserUploadParse(c *fasthttp.RequestCtx) {
	//todo: parse user upload and detect
	log.Println("todo: not implemented, upload")
}

func ShowSnip(c *fasthttp.RequestCtx) {
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
	case "showVerify":
		//todo: render verify html and output
		log.Println("verifying")
	default:
		//todo: really show snippet
		log.Println(tmpvar)
	}
}


func DeleteSnip(c *fasthttp.RequestCtx) {
	// todo: remove to use fasthttp as replace
	log.Println("todo: not implemented, delete")
}

func ShowVerifyCAPT(c *fasthttp.RequestCtx) {
	// todo: remove to use fasthttp as replace
	fmt.Println("todo: not implemented, verify")
}
