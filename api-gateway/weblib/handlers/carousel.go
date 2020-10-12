// Package handlers ..
/*
 * @Descripttion:处理函数
 * @Author: congz
 * @Date: 2020-09-20 11:33:37
 * @LastEditors: congz
 * @LastEditTime: 2020-10-12 21:44:25
 */
package handlers

import (
	"api-gateway/services"
	"context"

	"github.com/gin-gonic/gin"
)

//GetCarouselsList 使用rpc获取轮播图列表
func GetCarouselsList(ginCtx *gin.Context) {
	var carouselReq services.CarouselRequest
	PanicIfError(ginCtx.Bind(&carouselReq))
	//从gin.keys取出服务实例
	otherService := ginCtx.Keys["otherService"].(services.OtherService)
	//调用服务端的函数
	carouselRes, _ := otherService.GetCarouselsList(context.Background(), &carouselReq)
	ginCtx.JSON(200, gin.H{"data": gin.H{"carousel": carouselRes.CarouselsList, "count": carouselRes.Count}})
}

//GetCarousel 使用rpc获取轮播图
func GetCarousel(ginCtx *gin.Context) {
	var carouselReq services.CarouselRequest
	PanicIfError(ginCtx.BindUri(&carouselReq))
	//从gin.keys取出服务实例
	otherService := ginCtx.Keys["otherService"].(services.OtherService)
	carouselRes, _ := otherService.GetCarousel(context.Background(), &carouselReq)
	ginCtx.JSON(200, gin.H{"data": carouselRes.CarouselDetail})
}

//UpdateCarousel 使用rpc更新轮播图数据
func UpdateCarousel(ginCtx *gin.Context) {
	var carouselReq services.CarouselRequest
	PanicIfError(ginCtx.Bind(&carouselReq))
	//从gin.keys取出服务实例
	otherService := ginCtx.Keys["otherService"].(services.OtherService)
	carouselRes, _ := otherService.UpdateCarousel(context.Background(), &carouselReq)
	ginCtx.JSON(200, gin.H{"data": carouselRes.CarouselDetail})
}
