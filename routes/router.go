package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/planet-i/ginblog/common"
	"github.com/planet-i/ginblog/controller/article"
	"github.com/planet-i/ginblog/controller/category"
	"github.com/planet-i/ginblog/controller/user"
	"github.com/planet-i/ginblog/middleware"
	"github.com/planet-i/ginblog/utils"
)

//路由的入口文件
func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New() //初始化路由  Default 默认加了中间件
	r.Use(middleware.Log())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		//用户模块的路由接口
		auth.PUT("user/:id", user.EditUser)
		auth.DELETE("user/:id", user.DeleteUser) //去user的模型里写数据库操作的方法
		//分类模块的路由接口
		auth.POST("category/add", category.AddCategory)
		auth.PUT("category/:id", category.EditCategory)
		auth.DELETE("category/:id", category.DeleteCategory)
		//文章模块的路由接口
		auth.POST("article/add", article.AddArticle)
		auth.PUT("article/:id", article.EditArticle)
		auth.DELETE("article/:id", article.DeleteArticle)
		//上传文件
		auth.POST("upload", common.UpLoad)
	}
	router := r.Group("api/v1")
	{
		router.POST("user/add", user.AddUser)
		router.GET("users", user.GetUsers)
		router.GET("categories", category.GetCategories)
		router.GET("articles", article.GetArticles)             //查询所有文章
		router.GET("article/list/:cid", article.GetCateArticle) //查询分类下的所有文章
		router.GET("article/info/:id", article.GetArticleInfo)  //查询单个文章
		router.POST("login", user.Login)
	}
	r.Run(utils.HttpPort)
}
