package webserv

import (
	"encoding/base64"
	"encoding/binary"
	"github.com/pb-go/pb-go/config"
	"github.com/pb-go/pb-go/content_tools"
	"github.com/pb-go/pb-go/databaseop"
	"github.com/pb-go/pb-go/templates"
	"github.com/pb-go/pb-go/utils"
	_ "github.com/pb-go/pb-go/statik"
	"github.com/rakyll/statik/fs"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"
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
	// init obj
	var userForm = databaseop.UserData{}
	// first parse user form
	userPwd := c.FormValue("p")
	userExpire := int(binary.BigEndian.Uint16(c.FormValue("e")))
	userData := c.FormValue("d")
	// calculate expire
	// if recaptcha enabled, set to 5min expires first,
	if config.ServConf.Recaptcha.Enable {
		userForm.ExpireAt = primitive.NewDateTimeFromTime(time.Now().Add(5 * time.Minute))
	}
	primitive.NewDateTimeFromTime(time.Now().Add(userExpire * time.Hour))



	// then encrypt
	// then return recaptcha url, set id param in url using rawurl_b64.
	// else, set to 24hrs, then publish.
	// return publish url instead

}

func setShowSnipRenderData(userdt *databaseop.UserData, ctx *fasthttp.RequestCtx, israw bool) {
	decres, err := utils.DecryptData(userdt.Data.Data, ctx.FormValue("p"))
	if err != nil {
		ctx.SetStatusCode(http.StatusForbidden)
		return
	}
	ctx.SetStatusCode(http.StatusOK)
	if israw {
		ctx.SetContentType("text/plain")
		ctx.SetBody(decres)
		return
	} else {
		ctx.SetContentType("text/html; charset=utf-8")
		ctx.SetBodyString(templates.ShowSnipPageRend(string(decres)))
		return
	}
}

func readFromEmbed(statikfs http.FileSystem, filenm string, c *fasthttp.RequestCtx){
	tempfd, err := fs.ReadFile(statikfs, filenm)
	if err != nil {
		c.SetStatusCode(http.StatusNotFound)
		return
	}
	c.SetStatusCode(http.StatusOK)
	c.SetBody(tempfd)
	return
}

func ShowSnip(c *fasthttp.RequestCtx) {
	tmpvar := c.UserValue("shortId")
	switch tmpvar {
	case nil:
		fallthrough
	case "index.html":
		c.SetContentType("text/html; charset=utf-8")
		readFromEmbed(STFS, "/index.html", c)
		return
	case "submit.html":
		c.SetContentType("text/html; charset=utf-8")
		c.SetStatusCode(http.StatusOK)
		c.SetBodyString(templates.ShowSubmitPage())
		return
	case "favicon.ico":
		c.SetContentType("image/vnd.microsoft.icon")
		readFromEmbed(STFS, "/favicon.ico", c)
		return
	case "showVerify":
		c.SetContentType("text/html; charset=utf-8")
		c.SetStatusCode(http.StatusOK)
		c.SetBodyString(templates.VerifyPageRend())
		return
	case "status":
		c.SetContentType("text/plain")
		c.SetStatusCode(http.StatusOK)
		c.SetBodyString(templates.ServStatus())
		return
	default:
		filter1 := bson.M{"shortId": tmpvar}
		readoutDta, err := databaseop.GlobalMDBC.ItemRead(filter1)
		if err != nil || readoutDta.WaitVerify {
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
 			setShowSnipRenderData(&readoutDta, c, rawRender)
 			if readoutDta.ReadThenBurn {
				_ = databaseop.GlobalMDBC.ItemDelete(filter1)
			}
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
	legalkey := utils.GetUTCTimeHash(config.ServConf.Security.Master_key)
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
				{"expireAt", primitive.NewDateTimeFromTime(time.Now().Add(24 * time.Hour))},
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
