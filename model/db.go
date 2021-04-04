//db 的入口文件，配置 数据库连接参数
package model

import (
	"fmt"
	"time"

	"github.com/planet-i/ginblog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func initDB() {
	fmt.Println("db initing")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	) //处理连接文件 %s对应下面的变量
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
	}
	//db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
	//	DisableForeignKeyConstraintWhenMigrating: true,// 外键约束
	//	SkipDefaultTransaction: true, // 禁用默认事务（提高运行速度）
	//	NamingStrategy: schema.NamingStrategy{
	//	SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
	//	},
	//})// 处理error

	//配置 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, _ := db.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second) //虚拟一个连接池，需要使用的时候再连接   保证它不会大于框架连接时间
	DB = db
}

// GORM 使用结构体名的 蛇形命名 作为表名。对于结构体 User，根据约定，其表名为 users
func createdTable(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Article{})
	db.AutoMigrate(&Category{})
}

func init() {
	initDB()
	createdTable(DB)
}
