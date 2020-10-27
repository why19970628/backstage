// Package handlers ..
/*
 * @Descripttion:处理函数
 * @Author: congz
 * @Date: 2020-09-20 11:33:37
 * @LastEditors: congz
 * @LastEditTime: 2020-10-27 13:35:50
 */
package handlers

import (
	"api-gateway/pkg/e"
	"api-gateway/pkg/util"
	"api-gateway/services"
	"context"

	"github.com/gin-gonic/gin"
)

//AdminLogin 使用rpc进行管理员登录
func AdminLogin(ginCtx *gin.Context) {
	var adminReq services.AdminRequest
	PanicIfUserError(ginCtx.Bind(&adminReq))
	//从gin.keys取出服务实例
	userService := ginCtx.Keys["userService"].(services.UserService)
	//调用服务端的函数
	adminRes, _ := userService.AdminLogin(context.Background(), &adminReq)
	token, err := util.GenerateToken()
	if err != nil {
		adminRes.Code = e.ERROR_AUTH_TOKEN
	}
	ginCtx.JSON(200, gin.H{"code": adminRes.Code, "msg": e.GetMsg(adminRes.Code), "data": gin.H{"admin": adminRes.AdminDetail, "token": token}})
}
