// Package serviceimpl ...
/*
 * @Descripttion:商品服务
 * @Author: congz
 * @Date: 2020-09-15 10:57:26
 * @LastEditors: congz
 * @LastEditTime: 2020-09-21 16:21:08
 */
package serviceimpl

import (
	"context"
	"other/model"
	"other/services"
)

//BuildCarousel 序列化轮播图
func BuildCarousel(item model.Carousel) *services.CarouselModel {
	carouselModel := services.CarouselModel{
		ID:        uint32(item.ID),
		ProductID: item.ProductID,
		ImgPath:   item.ImgPath,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return &carouselModel
}

//OtherService 其他服务
type OtherService struct {
}

//GetCarouselsList 实现其他服务接口 获取轮播图列表
func (*OtherService) GetCarouselsList(context context.Context, req *services.CarouselRequest, res *services.CarouselsListResponse) error {
	carouselData := []model.Carousel{}
	model.DB.Find(&carouselData)
	carouselRes := []*services.CarouselModel{}
	for _, item := range carouselData {
		carouselRes = append(carouselRes, BuildCarousel(item))
	}
	res.CarouselsList = carouselRes
	return nil
}

//UpdateCarousel 实现其他服务接口 修改轮播图
func (*OtherService) UpdateCarousel(context context.Context, req *services.CarouselRequest, res *services.CarouselDetailResponse) error {
	return nil
}
