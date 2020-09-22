// Package serviceimpl ...
/*
 * @Descripttion:商品服务
 * @Author: congz
 * @Date: 2020-09-15 10:57:26
 * @LastEditors: congz
 * @LastEditTime: 2020-09-22 21:24:52
 */
package serviceimpl

import (
	"context"
	"other/model"
	"other/services"
)

//BuildNotice 序列化公告
func BuildNotice(item model.Notice) *services.NoticeModel {
	noticeModel := services.NoticeModel{
		ID:        uint32(item.ID),
		Text:      item.Text,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return &noticeModel
}

//GetNotice 实现其他服务接口 获取公告
func (*OtherService) GetNotice(ctx context.Context, req *services.NoticeRequest, res *services.NoticeDetailResponse) error {
	noticeData := model.Notice{}
	err := model.DB.First(&noticeData, req.NoticeID).Error
	if err != nil {
		return err
	}
	noticeRes := BuildNotice(noticeData)
	res.NoticeDetail = noticeRes
	return nil
}
