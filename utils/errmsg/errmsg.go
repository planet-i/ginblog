package errmsg

//jwt验证  json web token
//声明一系列常量，来写状态码
const (
	SUCCESS = 200
	ERROR   = 500
	// code= 1000...     用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_NOT_EXIST  = 1004 //用户写的token不存在
	ERROR_TOKNE_TIMEOUT    = 1005 //token超时
	EROOR_TOKEN_WRONG      = 1006 //用户写的token和我们验证出来的token是不一致的
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008
	// code = 2000....  文章模块的错误
	ERROR_ART_NOT_EXIST = 2001
	// code = 3000...   分类模块的错误
	ERROR_CATENAME_USED  = 3001
	ERROR_CATE_NOT_EXIST = 3002
)

//声明字典
var codeMsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    " 用户名已存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_NOT_EXIST:  "Token不存在，请重新登录",
	ERROR_TOKNE_TIMEOUT:    "Token已过期，请重新登录",
	EROOR_TOKEN_WRONG:      "Token不正确，请重新登录",
	ERROR_TOKEN_TYPE_WRONG: "Token格式错误，请重新登录",
	ERROR_CATENAME_USED:    "分类名已存在",
	ERROR_CATE_NOT_EXIST:   "该分类不存在",
	ERROR_ART_NOT_EXIST:    "文章不存在",
	ERROR_USER_NO_RIGHT:    "该用户无权限",
}

//处理返回错误信息
func GetErrmsg(code int) string {
	return codeMsg[code]
}
