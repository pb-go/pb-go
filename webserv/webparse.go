package webserv

import (
	"encoding/base64"
	"github.com/pb-go/pb-go/config"
	"github.com/pb-go/pb-go/content_tools"
	"github.com/pb-go/pb-go/databaseop"
	"github.com/pb-go/pb-go/templates"
	"github.com/pb-go/pb-go/utils"
	_ "github.com/pb-go/pb-go/statik"
	"github.com/rakyll/statik/fs"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

var (
	STFS http.FileSystem
)

func InitStatikFS(stfs *http.FileSystem){
	var err error
	*stfs, err = fs.New()
	if err != nil {
		log.Fatalln(err)
	}
}

func UserUploadParse(c *fasthttp.RequestCtx) {
	//todo: parse user upload and detect
	log.Println("todo: not implemented, upload")
}

func setShowSnipRenderData(userdt *databaseop.UserData, ctx *fasthttp.RequestCtx, israw bool){
	if israw {
		ctx.SetContentType("text/plain")
		//todo: decrypt data and output as binary
		ctx.SetBody(userdt.Data.Data)
		return
	} else {
		ctx.SetContentType("text/html; charset=utf-8")
		//todo: decrypt data and output as binary
		ctx.SetBodyString(templates.ShowSnipPageRend(string(userdt.Data.Data)))
		return
	}
}

func ShowSnip(c *fasthttp.RequestCtx) {
	var tempfd []byte
	tmpvar := c.UserValue("shortId")
	switch tmpvar {
	case nil:
		fallthrough
	case "index.html":
		tempfd, _ = fs.ReadFile(STFS, "/index.html")
		c.SetBody(tempfd)
		return
	case "submit.html":
		tempfd, _ = fs.ReadFile(STFS, "/submit.html")
		c.SetBody(tempfd)
		return
	case "favicon.ico":
		tempfd, _ = fs.ReadFile(STFS, "/favicon.ico")
		c.SetBody(tempfd)
		return
	case "showVerify":
		c.SetContentType("text/html")
		c.SetStatusCode(http.StatusOK)
		c.SetBodyString(templates.VerifyPageRend(config.ServConf.Recaptcha.Site_key))
		return
	default:
		filter1 := bson.M{"shortId": tmpvar}
		readoutDta, err := databaseop.GlobalMDBC.ItemRead(filter1)
		if err != nil {
			c.SetStatusCode(http.StatusBadRequest)
			return
 		} else {
 			var rawRender = string(c.FormValue("f")) == "raw"
 			if readoutDta.PwdIsSet {
 				uploadedpwd := c.FormValue("p")
 				hashedupdpwd := utils.GenBlake2B(uploadedpwd)
 				if hashedupdpwd != readoutDta.Password {
					c.SetStatusCode(http.StatusForbidden)
					return
				}
			}
			c.SetStatusCode(http.StatusOK)
 			setShowSnipRenderData(&readoutDta, c, rawRender)
 			return
		}
	}
}

func DeleteSnip(c *fasthttp.RequestCtx) {
	masterkey := string(c.Request.Header.Peek("X-Master-Key"))
	if masterkey == "" {
		c.SetStatusCode(http.StatusForbidden)
		return
	}
	legalkey := utils.GetUTCTimeHash()
	if masterkey == legalkey {
		curshortid := string(c.FormValue("id"))
		filter1 := bson.M{"shortId": curshortid}
		err := databaseop.GlobalMDBC.ItemDelete(filter1)
		if err != nil {
			c.SetStatusCode(http.StatusBadRequest)
			return
 		} else {
 			c.SetStatusCode(http.StatusAccepted)
 			return
		}
	} else {
		c.SetStatusCode(http.StatusForbidden)
		return
	}
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
	rmtIPhd := string(c.Request.Header.Peek("X-Real-IP"))
	if rmtIPhd == "" {
		c.SetStatusCode(http.StatusBadGateway)
		return
	}
	res, err := content_tools.VerifyRecaptchaResp(string(c.FormValue("g-recaptcha-response")), rmtIPhd)
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
			log.Println(err)
			c.SetStatusCode(http.StatusGone)
			return
 		} else {
 			c.SetStatusCode(http.StatusOK)
 			c.SetContentType("text/plain")
 			c.SetBodyString("Verification Passed. Go to https://"+ config.ServConf.Network.Host + "/" + current_snipid + " to see your paste.")
 			return
		}
	}
}
