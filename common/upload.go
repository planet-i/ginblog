package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/planet-i/ginblog/model"
	"github.com/planet-i/ginblog/utils/errmsg"
)

func UpLoad(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")

	fileSize := fileHeader.Size

	url, code := model.UpLoadFile(file, fileSize)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrmsg(code),
		"url":     url,
	})
}
