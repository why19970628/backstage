// Package serviceimpl ...
/*
 * @Descripttion:商品服务
 * @Author: congz
 * @Date: 2020-09-15 10:57:26
 * @LastEditors: congz
 * @LastEditTime: 2020-09-20 16:10:15
 */
package serviceimpl

import (
	"context"
	"product/model"
	"product/services"
)

//BuildProduct 序列化商品
func BuildProduct(item model.Product) *services.ProductModel {
	productModel := services.ProductModel{
		ID:            uint32(item.ID),
		Name:          item.Name,
		CategoryID:    item.CategoryID,
		Title:         item.Title,
		Info:          item.Info,
		ImgPath:       item.ImgPath,
		Price:         item.Price,
		DiscountPrice: item.DiscountPrice,
		CreatedAt:     item.CreatedAt.Unix(),
		UpdatedAt:     item.UpdatedAt.Unix(),
	}
	return &productModel
}

//ProductService 商品服务
type ProductService struct {
}

//GetProductsList 实现商品服务接口 获取商品列表
func (*ProductService) GetProductsList(ctx context.Context, req *services.ProductRequest, res *services.ProductsListResponse) error {
	if req.Limit == 0 {
		req.Limit = 6
	}
	//在数据库查找值
	productData := []model.Product{}
	model.DB.Offset(req.Start).Limit(req.Limit).Find(&productData)
	//序类化商品列表
	productRes := []*services.ProductModel{}
	for _, item := range productData {
		productRes = append(productRes, BuildProduct(item))
	}
	//序列化后的结果赋给response
	res.ProductList = productRes
	return nil
}

//GetProductDetail 实现商品服务接口 获取商品详情
func (*ProductService) GetProductDetail(ctx context.Context, req *services.ProductRequest, res *services.ProductDetailResponse) error {
	//在数据库查找值
	productData := model.Product{}
	model.DB.First(&productData, req.ProductId)
	//序类化商品
	productRes := BuildProduct(productData)
	//序列化后的结果赋给response
	res.ProductDetail = productRes
	return nil
}
