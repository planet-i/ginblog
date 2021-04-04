package model

import (
	"github.com/planet-i/ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"` //分类
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

//新增文章
func AddArticle(data *Article) int {
	err := DB.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

// 删除文章
func DeleteArticle(id int) int {
	var art Article
	err = DB.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 编辑文章
func EditArticle(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err = DB.Model(&art).Where("id = ?", id).Updates(maps).Error //编辑不存在的id也会成功
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询文章列表
func GetArticles(pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var total int64
	err = DB.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCESS, total
}

//查询单个文章
func GetArticleInfo(id int) (Article, int) {
	var art Article
	err := DB.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCESS
}

//查询分类下所有文章
func GetCateArticle(cid int, pageSize int, pageNum int) ([]Article, int, int64) {
	var cateArtList []Article
	var total int64
	err = DB.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", cid).Find(&cateArtList).Count(&total).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return cateArtList, errmsg.SUCCESS, total
}
