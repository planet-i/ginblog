package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/planet-i/ginblog/middleware"
	"github.com/planet-i/ginblog/model"
	"github.com/planet-i/ginblog/utils/errmsg"
)

//为什么会有前台登陆和后台登陆两个界面呢，登陆的时候判断角色码然后进入不同界面不就可以了吗
//后台登陆
func Login(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)
	var token string
	var code int
	data, code = model.CheckLogin(data.Username, data.Password)
	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data.Username,
		"id":      data.ID,
		"message": errmsg.GetErrmsg(code),
		"token":   token,
	})
}

// 前台登录
func LoginFront(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	var code int

	data, code = model.CheckLoginFront(data.Username, data.Password)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data.Username,
		"id":      data.ID,
		"message": errmsg.GetErrmsg(code),
	})
}

type UpToken struct {
	Token string `json:"token"`
}

// 验证token
func CheckToken(c *gin.Context) {
	var Token UpToken
	_ = c.ShouldBindJSON(&Token)

	_, code = middleware.CheckToken(Token.Token)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrmsg(code),
	})
}
