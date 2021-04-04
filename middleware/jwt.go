package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/planet-i/ginblog/utils"
	"github.com/planet-i/ginblog/utils/errmsg"
)

var JwtKey = []byte(utils.JwtKey) //用来生产token命令    在config里添加做网站参数配置，setting里面设置

//用结构体接收参数    与用户模型里的一致，嵌套一个jwt自带的结构体
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var code int

//生成token
func SetToken(username string) (string, int) {
	//给个区间有效期
	expireTime := time.Now().Add(10 * time.Hour)
	SetClaims := MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ginblog",
		},
	}
	//把结构体赋值进去
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims) //加盐解析
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return " ", errmsg.ERROR
	}
	//需要一个方法，传参数（签发的标准方法，
	return token, errmsg.SUCCESS
}

//验证token
func CheckToken(token string) (*MyClaims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, _ := setToken.Claims.(*MyClaims); setToken.Valid {
		return key, errmsg.SUCCESS //解析token 比对，正确/错误
	} else {
		return nil, errmsg.ERROR
	}
}

//jwt中间件  控制验证的token写入router中
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHerder := c.Request.Header.Get("Authorization")
		if tokenHerder == "" {
			code = errmsg.ERROR_TOKEN_NOT_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrmsg(code),
			})
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHerder, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrmsg(code),
			})
			c.Abort()
			return
		}
		key, tCode := CheckToken(checkToken[1])
		if tCode == errmsg.ERROR {
			code = errmsg.EROOR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrmsg(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ERROR_TOKNE_TIMEOUT
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrmsg(code),
			})
			c.Abort()
			return
		}
		c.Set("username", key.Username)
		c.Next()
	}

}

//规范
//生成一个固定格式
