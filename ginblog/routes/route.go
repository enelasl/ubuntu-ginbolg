package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/utils"

	"github.com/gin-gonic/gin"
)

func IninRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default() //初始化路由

	router := r.Group("api/v1") //设置路由组
	{
		//User
		router.POST("user/add", v1.AddUser)
		router.GET("user/add", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DelectUser)

		//Category
		router.GET("admin/category", v1.GetCate)
		router.POST("category/add", v1.AddCategory)
		router.PUT("category/:id", v1.EditCate)
		router.DELETE("category/:id", v1.DeleteCate)
		//Article

	}

	r.Run(utils.HttpPort)
}
