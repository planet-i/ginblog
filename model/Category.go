package model

import (
	"github.com/planet-i/ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

//查询分类是否存在
func CheckCategory(name string) int {
	var cate Category
	DB.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

//新增分类
func AddCategory(data *Category) int {
	err := DB.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

// 删除分类        如果不用gorm.model的话就是硬删除
func DeleteCategory(id int) int {
	var cate Category
	err = DB.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 编辑分类
func EditCategory(id int, data *Category) (int, int64) {
	var cate Category
	var total int64
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = DB.Model(&cate).Where("id = ?", id).Updates(maps).Count(&total).Error //编辑不存在的id也会成功
	if err != nil {
		return errmsg.ERROR, 0
	}
	return errmsg.SUCCESS, total //gorm不识别大小写？
}

// 查询单个分类信息
func GetCateInfo(id int) (Category, int) {
	var cate Category
	DB.Where("id = ?", id).First(&cate)
	return cate, errmsg.SUCCESS
}

//获取分类列表
func GetCategories(pageSize int, pageNum int) ([]Category, int64) {
	var cate []Category
	var total int64
	err = DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cate, total
}
