// Package weblib ..
/*
 * @Descripttion:新建gin路由
 * @Author: congz
 * @Date: 2020-09-20 11:33:37
 * @LastEditors: congz
 * @LastEditTime: 2020-10-12 22:33:04
 */
package weblib

import (
	"api-gateway/weblib/handlers"
	"api-gateway/weblib/middlewares"

	"github.com/gin-gonic/gin"
)

//NewRouter 新建gin路由
func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	//使用中间件，接收服务调用实例
	ginRouter.Use(middlewares.Cors(), middlewares.InitMiddleware(service), middlewares.ErrorMiddleware())
	v1 := ginRouter.Group("/api/v1")
	{
		//需要登录验证
		authed := v1.Group("/")
		{
			authed.GET("/products", handlers.GetProductsList)
			authed.GET("/products/:product_id", handlers.GetProductDetail)
		}
		//轮播图服务
		v1.GET("/carousels", handlers.GetCarouselsList)
		v1.GET("/carousels/:carousel_id", handlers.GetCarousel)
		v1.PUT("/carousels", handlers.UpdateCarousel)
		//公告服务
		v1.POST("/notices", handlers.CreateNotice)
		v1.GET("/notices/:notice_id", handlers.GetNotice)
		v1.PUT("/notices", handlers.UpdateNotice)
		v1.DELETE("/notices", handlers.DeleteNotice)
		//分类服务
		v1.POST("/categories", handlers.CreateCategory)
		v1.GET("/categories", handlers.GetCategoriesList)
		v1.PUT("/categories", handlers.UpdateCategory)
		v1.DELETE("/categories", handlers.DeleteCategory)
		//用户服务
		v1.POST("/admins", handlers.AdminLogin)
		v1.GET("/users", handlers.GetUsersList)
		v1.GET("/users/:user_id", handlers.GetUser)
	}
	return ginRouter
}
