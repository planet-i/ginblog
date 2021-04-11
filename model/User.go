package model

import (
	"log"

	"github.com/planet-i/ginblog/utils/errmsg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//全局可以引用的结构体
type User struct {
	gorm.Model
	//gorm提供的四个参数，ID，创建时间、更新时间、删除时间
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	//json跟前后端数据的交互有关，上传图片用form表单形式
	Password string `gorm:"type:varchar(256);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
	//1 管理员，2 用户
	// 头像 Avater string
}

//json格式，根据路由的传输数据来绑定 gin框架给了json绑定的方式
//很多东西在前端填表单的时候就可以验证，不让它提交即可，比后端灵活
//这里只用GIN针对数据库操作的验证

//与接口保持一致 ，相当于对数据库的操作

//查询用户是否存在
func CheckUser(name string) int {
	//引用user，才能在里面取值
	var users User
	DB.Select("id").Where("username = ?", name).First(&users) //DB
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// 更新查询
func CheckUpUser(id int, name string) (code int) {
	var user User
	DB.Select("id, username").Where("username = ?", name).First(&user)
	if user.ID == uint(id) {
		return errmsg.SUCCESS
	}
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCESS
}

//新增用户  传递模型   结构体是引用型类型   函数里作为入参传递的话应该传指针  返回code
func AddUser(data *User) int {
	data.Password = ScryptPw(data.Password)
	err := DB.Create(&data).Error // DB
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

//查询单个用户
func GetUser(id int) (User, int) {
	var user User
	err = DB.Where("ID = ?", id).First(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCESS
}

//有列表的东西，肯定会涉及分页   返回User模型的切片
//获取用户列表
func GetUsers(username string, pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64
	if username == "" {
		DB.Select("id,username,role").Where(
			"username LIKE ?", username+"%",
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
		DB.Model(&users).Where(
			"username LIKE ?", username+"%",
		).Count(&total)
		return users, total
	}
	DB.Select("id,username,role").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
	DB.Model(&users).Count(&total)
	// if err != nil{
	// 	return users, 0
	// }
	return users, total
}

// 编辑用户信息，忽略密码的修改
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = DB.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 修改密码
func ChangePassword(id int, data *User) int {
	//var user User
	//var maps = make(map[string]interface{})
	//maps["password"] = data.Password
	//存进去的数据加不加密要怎么处理
	data.Password = ScryptPw(data.Password)
	err = DB.Select("password").Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除用户
func DeleteUser(id int) int {
	var user User
	err = DB.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// // 密码加密
// func ScryptPws(password string) string {
// 	const KeyLen = 10
// 	salt := make([]byte, 8)
// 	salt = []byte{12, 32, 4, 6, 33, 22, 222, 11}
// 	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
// 	用包内的另一个方法去生成密码
//	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fpw := base64.StdEncoding.EncodeToString(HashPw)
// 	return fpw
// }

// 密码加密&权限控制
// func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
// 	u.Password = ScryptPw(u.Password)
// 	u.Role = 2
// 	return nil
// }

// func (u *User) BeforeUpdate(_ *gorm.DB) (err error) {
// 	u.Password = ScryptPw(u.Password)
// 	return nil
// }

// 生成密码
func ScryptPw(password string) string {
	const cost = 10
	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}
	return string(HashPw)
}

//后台登陆验证
func CheckLogin(username string, password string) (User, int) {
	var user User
	var PasswordErr error
	DB.Where("username = ?", username).First(&user)
	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	// if ScryptPw(password) != user.Password {
	// 	return user, errmsg.ERROR_PASSWORD_WRONG
	// }
	if PasswordErr != nil {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return user, errmsg.ERROR_USER_NO_RIGHT
	}
	return user, errmsg.SUCCESS
}

// 前台登录
func CheckLoginFront(username string, password string) (User, int) {
	var user User
	var PasswordErr error

	DB.Where("username = ?", username).First(&user)

	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	if PasswordErr != nil {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	return user, errmsg.SUCCESS
}
