// Package handlers ..
/*
 * @Descripttion:商品图片处理函数
 * @Author: congz
 * @Date: 2020-09-20 11:33:37
 * @LastEditors: congz
 * @LastEditTime: 2020-10-14 20:18:50
 */
package handlers

import (
	"api-gateway/services"
	"context"

	"github.com/gin-gonic/gin"
)

//CreateProductImg 使用rpc创建商品图片
func CreateProductImg(ginCtx *gin.Context) {
	var productImgReq services.ProductImgRequest
	PanicIfError(ginCtx.Bind(&productImgReq))
	//从gin.keys取出服务实例
	productService := ginCtx.Keys["productService"].(services.ProductService)
	//调用服务端的函数
	productImgRes, _ := productService.CreateProductImg(context.Background(), &productImgReq)
	ginCtx.JSON(200, gin.H{"data": productImgRes.ProductImgDetail})
}

//GetProductImgsList 使用rpc获取商品图片列表
func GetProductImgsList(ginCtx *gin.Context) {
	var productImgReq services.ProductImgRequest
	PanicIfError(ginCtx.Bind(&productImgReq))
	//从gin.keys取出服务实例
	productService := ginCtx.Keys["productService"].(services.ProductService)
	//调用服务端的函数
	productImgRes, _ := productService.GetProductImgsList(context.Background(), &productImgReq)
	ginCtx.JSON(200, gin.H{"data": gin.H{"product_img": productImgRes.ProductImgsList, "count": productImgRes.Count}})
}

//UpdateProductImg 使用rpc更新商品图片
func UpdateProductImg(ginCtx *gin.Context) {
	var productImgReq services.ProductImgRequest
	PanicIfError(ginCtx.Bind(&productImgReq))
	//从gin.keys取出服务实例
	productService := ginCtx.Keys["productService"].(services.ProductService)
	//调用服务端的函数
	productImgRes, _ := productService.UpdateProductImg(context.Background(), &productImgReq)
	ginCtx.JSON(200, gin.H{"data": productImgRes.ProductImgDetail})
}

//DeleteProductImg 使用rpc删除商品图片
func DeleteProductImg(ginCtx *gin.Context) {
	var productImgReq services.ProductImgRequest
	PanicIfError(ginCtx.Bind(&productImgReq))
	//从gin.keys取出服务实例
	productService := ginCtx.Keys["productService"].(services.ProductService)
	//调用服务端的函数
	productImgRes, _ := productService.DeleteProductImg(context.Background(), &productImgReq)
	ginCtx.JSON(200, gin.H{"data": productImgRes.ProductImgDetail})
}
