// Package handlers ..
/*
 * @Descripttion:处理函数
 * @Author: congz
 * @Date: 2020-09-20 11:33:37
 * @LastEditors: congz
 * @LastEditTime: 2020-10-12 13:55:35
 */
package handlers

import (
	"api-gateway/services"
	"context"

	"github.com/gin-gonic/gin"
)

//GetUsersList 使用rpc获取用户列表
func GetUsersList(ginCtx *gin.Context) {
	var userReq services.UserRequest
	PanicIfError(ginCtx.Bind(&userReq))
	//从gin.keys取出服务实例
	userService := ginCtx.Keys["userService"].(services.UserService)
	//调用服务端的函数
	userRes, _ := userService.GetUsersList(context.Background(), &userReq)
	ginCtx.JSON(200, gin.H{"data": gin.H{"user": userRes.UsersList, "count": userRes.Count}})
}

//GetUser 使用rpc获取用户详情
func GetUser(ginCtx *gin.Context) {
	var userReq services.UserRequest
	PanicIfError(ginCtx.BindUri(&userReq))
	//从gin.keys取出服务实例
	userService := ginCtx.Keys["userService"].(services.UserService)
	//调用服务端的函数
	userRes, _ := userService.GetUser(context.Background(), &userReq)
	ginCtx.JSON(200, gin.H{"data": userRes.UserDetail})
}
