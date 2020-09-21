// Package weblib ..
/*
 * @Descripttion:新建gin路由
 * @Author: congz
 * @Date: 2020-09-20 11:33:37
 * @LastEditors: congz
 * @LastEditTime: 2020-09-21 15:51:22
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
	ginRouter.Use(middlewares.InitMiddleware(service), middlewares.ErrorMiddleware())
	v1 := ginRouter.Group("/api/v1")
	{
		//需要登录验证
		authed := v1.Group("/")
		{
			authed.Use(middlewares.JWT())
			authed.POST("/products", handlers.GetProductsList)
			authed.GET("/products/:product_id", handlers.GetProductDetail)
		}
		v1.GET("/carousels", handlers.GetCarouselsList)
		v1.POST("/carousels", handlers.UpdateCarousel)

		v1.POST("/admins", handlers.AdminLogin)
	}
	return ginRouter
}
