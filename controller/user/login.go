package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/planet-i/ginblog/middleware"
	"github.com/planet-i/ginblog/model"
	"github.com/planet-i/ginblog/utils/errmsg"
)

func Login(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)
	var token string
	code := model.CheckLogin(data.Username, data.Password)
	if code == errmsg.SUCCESS {
		token, _ = middleware.SetToken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrmsg(code),
		"token":   token,
	})
}
