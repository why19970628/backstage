// Package serviceimpl ...
/*
 * @Descripttion:商品服务
 * @Author: congz
 * @Date: 2020-09-15 10:57:26
 * @LastEditors: congz
 * @LastEditTime: 2020-09-21 15:28:09
 */
package serviceimpl

import (
	"context"
	"user/model"
	"user/services"

	"github.com/jinzhu/gorm"
)

//BuildAdmin 序列化管理员
func BuildAdmin(item model.Admin) *services.AdminModel {
	adminModel := services.AdminModel{
		ID:        uint32(item.ID),
		UserName:  item.UserName,
		Avatar:    item.Avatar,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return &adminModel
}

//UserService 用户服务
type UserService struct {
}

//AdminLogin 实现用户服务接口 管理员登录
func (*UserService) AdminLogin(ctx context.Context, req *services.AdminRequest, res *services.AdminResponse) error {
	var admin model.Admin
	res.Code = 200
	if err := model.DB.Where("user_name = ?", req.UserName).First(&admin).Error; err != nil {
		//如果查询不到，返回相应错误
		if gorm.IsRecordNotFoundError(err) {
			res.Code = 10003
			return nil
		}
		res.Code = 30001
		return nil
	}
	if admin.CheckPassword(req.Password) == false {
		res.Code = 10004
		return nil
	}
	adminRes := BuildAdmin(admin)
	res.AdminDetail = adminRes
	return nil
}
