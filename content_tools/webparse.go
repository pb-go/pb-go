package content_tools

import (
	"encoding/base64"
	"github.com/kmahyyg/pb-go/config"
	"github.com/kmahyyg/pb-go/databaseop"
	"github.com/kmahyyg/pb-go/templates"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"github.com/valyala/fasthttp"
	"net/http"
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

func StartVerifyCAPT(c *fasthttp.RequestCtx) {
	if !config.ServConf.Recaptcha.Enable {
		c.SetStatusCode(http.StatusForbidden)
		return
	}
	var formsnipid []byte
	_ , err := base64.RawURLEncoding.Decode(formsnipid, c.FormValue("snipid"))
	current_snipid := string(formsnipid)
	if err != nil || current_snipid == "" {
		c.SetStatusCode(http.StatusBadRequest)
		return
	}
	res, err := VerifyRecaptchaResp(string(c.FormValue("g-recaptcha-response")), c.RemoteIP().String())
	if err != nil || res == false {
		c.SetStatusCode(http.StatusForbidden)
		return
	}
	if res == true {
		filter1 := bson.M{"shortId": formsnipid}
		update1 := bson.D{
			{"$set", bson.D {
				{"waitVerify", false},
			}},
		}
		err = databaseop.GlobalMDBC.ItemUpdate(filter1, update1)
		if err != nil {
			log.Println("#############NOTE: ILLEGAL REQUEST!#############")
			log.Println(err)
			c.SetStatusCode(http.StatusBadRequest)
			return
 		} else {
 			c.SetStatusCode(http.StatusOK)
 			c.SetContentType("text/html")
 			c.Set
		}
	}
}
