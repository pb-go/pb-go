package content_tools

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UserUploadParse(c *gin.Context) {
	// todo: remove to use fasthttp as replace
	panic("todo: not implemented")
}

func DefaultHand(c *gin.Context) {
	// todo: remove to use fasthttp as replace
	snipid := c.Param("shortId")
	log.Println([]byte(snipid))
	switch snipid {
	case "":

	}
	c.JSON(http.StatusOK, gin.H{
		"message": "boom",
	})
}

func ShowSnip(c *gin.Context) {
	// todo: remove to use fasthttp as replace
}


func DeleteSnip(c *gin.Context) {
	// todo: remove to use fasthttp as replace
	panic("todo: not implemented")
}

func VerifyCAPT(c *gin.Context) {
	// todo: remove to use fasthttp as replace
	panic("todo: not implemented")
}
