// Package serviceimpl ...
/*
 * @Descripttion:商品服务
 * @Author: congz
 * @Date: 2020-09-15 10:57:26
 * @LastEditors: congz
 * @LastEditTime: 2020-10-14 12:42:53
 */
package serviceimpl

import (
	"context"
	"product/model"
	"product/services"
)

//CreateProduct 实现商品服务接口 创建商品
func (*ProductService) CreateProduct(ctx context.Context, req *services.ProductRequest, res *services.ProductDetailResponse) error {
	//在数据库查找值
	productData := model.Product{
		CategoryID:    req.CategoryID,
		Name:          req.Name,
		Title:         req.Title,
		Info:          req.Info,
		ImgPath:       req.ImgPath,
		Price:         req.Price,
		DiscountPrice: req.DiscountPrice,
	}
	err := model.DB.Create(&productData).Error
	if err != nil {
		return err
	}
	//序列化后的结果赋给response
	res.ProductDetail = BuildProduct(productData)
	return nil
}

//GetProductsList 实现商品服务接口 获取商品列表
func (*ProductService) GetProductsList(ctx context.Context, req *services.ProductRequest, res *services.ProductsListResponse) error {
	if req.Limit == 0 {
		req.Limit = 6
	}
	//在数据库查找值
	productData := []model.Product{}
	var count uint32
	if req.CategoryID == 0 {
		err := model.DB.Offset(req.Start).Limit(req.Limit).Find(&productData).Error
		if err != nil {
			return err
		}
		err = model.DB.Model(&model.Product{}).Count(&count).Error
		if err != nil {
			return err
		}
	} else {
		err := model.DB.Offset(req.Start).Limit(req.Limit).Where("category_id=?", req.CategoryID).Find(&productData).Error
		if err != nil {
			return err
		}
		err = model.DB.Model(&model.Product{}).Where("category_id=?", req.CategoryID).Count(&count).Error
		if err != nil {
			return err
		}
	}

	//序类化商品列表
	productRes := []*services.ProductModel{}
	for _, item := range productData {
		productRes = append(productRes, BuildProduct(item))
	}
	//序列化后的结果赋给response
	res.ProductsList = productRes
	res.Count = count
	return nil
}

//GetProduct 实现商品服务接口 获取商品详情
func (*ProductService) GetProduct(ctx context.Context, req *services.ProductRequest, res *services.ProductDetailResponse) error {
	//在数据库查找值
	productData := model.Product{}
	err := model.DB.First(&productData, req.ProductID).Error
	if err != nil {
		return err
	}
	//序类化商品
	productRes := BuildProduct(productData)
	//序列化后的结果赋给response
	res.ProductDetail = productRes
	return nil
}

//UpdateProduct 实现商品服务接口 修改商品详情
func (*ProductService) UpdateProduct(ctx context.Context, req *services.ProductRequest, res *services.ProductDetailResponse) error {
	//在数据库查找值
	productData := model.Product{}
	err := model.DB.First(&productData, req.ProductID).Error
	if err != nil {
		return err
	}
	//将要更新的数据赋值给结构体
	productData.CategoryID = req.CategoryID
	productData.Name = req.Name
	productData.Title = req.Title
	productData.Info = req.Info
	productData.ImgPath = req.ImgPath
	productData.Price = req.Price
	productData.DiscountPrice = req.DiscountPrice
	//update
	err = model.DB.Save(&productData).Error
	if err != nil {
		return err
	}
	//序列化后的结果赋给response
	res.ProductDetail = BuildProduct(productData)
	return nil
}

//DeleteProduct 实现商品服务接口 删除商品
func (*ProductService) DeleteProduct(ctx context.Context, req *services.ProductRequest, res *services.ProductDetailResponse) error {
	//在数据库删除值
	err := model.DB.Delete(&model.Product{}, req.ProductID).Error
	if err != nil {
		return err
	}
	return nil
}
