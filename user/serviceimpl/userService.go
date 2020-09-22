// Package serviceimpl ...
/*
 * @Descripttion:商品服务
 * @Author: congz
 * @Date: 2020-09-15 10:57:26
 * @LastEditors: congz
 * @LastEditTime: 2020-09-22 21:51:35
 */
package serviceimpl

import (
	"context"
	"user/model"
	"user/services"
)

//BuildUser 序列化用户
func BuildUser(item model.User) *services.UserModel {
	userModel := services.UserModel{
		ID:        uint32(item.ID),
		UserName:  item.UserName,
		Avatar:    item.Avatar,
		Email:     item.Email,
		NickName:  item.Nickname,
		Status:    item.Status,
		Limit:     item.Limit,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return &userModel
}

//GetUsersList 实现用户服务接口 获取用户列表
func (*UserService) GetUsersList(ctx context.Context, req *services.UserRequest, res *services.UsersListResponse) error {
	if req.Limit == 0 {
		req.Limit = 6
	}
	//在数据库查找值
	userData := []model.User{}
	err := model.DB.Offset(req.Start).Limit(req.Limit).Find(&userData).Error
	if err != nil {
		return err
	}
	//序类化用户列表
	userRes := []*services.UserModel{}
	for _, item := range userData {
		userRes = append(userRes, BuildUser(item))
	}
	//序列化后的结果赋给response
	res.UsersList = userRes
	return nil
}
