package content_tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserUploadParse(c *gin.Context) {
	panic("todo: not implemented")
}

func ShowSnip(c *gin.Context) {
	//panic("todo: not implemented")
	c.JSON(http.StatusOK, gin.H{
		"message": "boom",
	})
}

func DeleteSnip(c *gin.Context) {
	panic("todo: not implemented")
}

func VerifyCAPT(c *gin.Context) {
	panic("todo: not implemented")
}
