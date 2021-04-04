package main

import (
	"fmt"

	"github.com/planet-i/ginblog/model"
)

func main() {
	//引用数据库
	model.DB.AutoMigrate(model.User{}) //自动迁移   对应生成model中的其他表
	model.DB.AutoMigrate(model.Article{})
	model.DB.AutoMigrate(model.Category{})
	initDBData()
}

func initDBData() {
	var admin model.User
	model.DB.Where("username = ?", "admin").First(&admin)
	fmt.Println(admin.Username)
	if admin.ID == 0 {
		admin.Username = "admin"
		admin.Password = model.ScryptPw("silin123")
		admin.Role = 1
		model.DB.Save(&admin)
	}

	// create user
	var user1 model.User
	model.DB.Where("username = ?", "user1").First(&user1)
	if user1.ID == 0 {
		user1.Username = "user1"
		user1.Password = model.ScryptPw("silin123")
		user1.Role = 2
		model.DB.Save(&user1)
	}

}
