// Package handlers ..
/*
 * @Descripttion:处理函数
 * @Author: congz
 * @Date: 2020-09-20 11:33:37
 * @LastEditors: congz
 * @LastEditTime: 2020-09-22 21:45:34
 */
package handlers

import (
	"api-gateway/services"
	"context"

	"github.com/gin-gonic/gin"
)

//GetProductsList 使用rpc获取商品列表
func GetProductsList(ginCtx *gin.Context) {
	var productReq services.ProductRequest
	PanicIfError(ginCtx.Bind(&productReq))
	//从gin.keys取出服务实例
	productService := ginCtx.Keys["productService"].(services.ProductService)
	//调用服务端的函数
	productRes, _ := productService.GetProductsList(context.Background(), &productReq)
	ginCtx.JSON(200, gin.H{"data": productRes.ProductsList})

}

//GetProductDetail 使用rpc获取商品详情
func GetProductDetail(ginCtx *gin.Context) {
	var productReq services.ProductRequest
	PanicIfError(ginCtx.BindUri(&productReq))
	//从gin.keys取出服务实例
	productService := ginCtx.Keys["productService"].(services.ProductService)
	productRes, _ := productService.GetProduct(context.Background(), &productReq)
	ginCtx.JSON(200, gin.H{"data": productRes.ProductDetail})
}
