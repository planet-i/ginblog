package user

import (
	"net/http"
	"strconv"

	"github.com/planet-i/ginblog/model"
	"github.com/planet-i/ginblog/utils/errmsg"
	"github.com/planet-i/ginblog/utils/validator"

	"github.com/gin-gonic/gin"
)

var code int

//相当于控制器，控制读写，在这里调用model里对数据库的操作，实现控制器的功能
//使用结构体方式把模型引进来

//添加用户
func AddUser(c *gin.Context) {
	//拿到用户名
	//对用户名进行检查
	var data model.User //引用结构体
	var msg string
	var validCode int
	_ = c.ShouldBindJSON(&data) //gin里面的绑定
	msg, validCode = validator.Validate(&data)
	if validCode != errmsg.SUCCESS {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  validCode,
				"message": msg,
			},
		)
		c.Abort()
		return
	}

	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.AddUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}
	//这里的200是网络传输里的，反映网络是否通畅
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrmsg(code),
		},
	)
	//role数据传输有问题，无论传什么都是2
}

//查询单个用户
func GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var maps = make(map[string]interface{})
	data, code := model.GetUser(id)
	maps["username"] = data.Username
	maps["role"] = data.Role
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   1,
			"message": errmsg.GetErrmsg(code),
		},
	)
}

//查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	data, total := model.GetUsers(username, pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrmsg(code),
	})
}

//编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.CheckUpUser(id, data.Username)
	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrmsg(code),
	})
}

// 修改密码
func ChangeUserPassword(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code = model.ChangePassword(id, &data)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrmsg(code),
		},
	)
}

//删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(
		http.StatusOK, gin.H{
			"ststus":  code,
			"message": errmsg.GetErrmsg(code),
		},
	)
}
