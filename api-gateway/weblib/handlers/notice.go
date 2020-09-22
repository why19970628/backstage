// Package handlers ..
/*
 * @Descripttion:处理函数
 * @Author: congz
 * @Date: 2020-09-20 11:33:37
 * @LastEditors: congz
 * @LastEditTime: 2020-09-22 21:32:04
 */
package handlers

import (
	"api-gateway/services"
	"context"

	"github.com/gin-gonic/gin"
)

//GetNotice 使用rpc获取公告
func GetNotice(ginCtx *gin.Context) {
	var noticeReq services.NoticeRequest
	PanicIfError(ginCtx.BindUri(&noticeReq))
	//从gin.keys取出服务实例
	otherService := ginCtx.Keys["otherService"].(services.OtherService)
	//调用服务端的函数
	noticeRes, _ := otherService.GetNotice(context.Background(), &noticeReq)
	ginCtx.JSON(200, gin.H{"data": noticeRes.NoticeDetail})

}
