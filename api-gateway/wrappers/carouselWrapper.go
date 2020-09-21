// Package wrappers ..
/*
 * @Descripttion:go-micro中间件
 * @Author: congz
 * @Date: 2020-09-15 22:22:40
 * @LastEditors: congz
 * @LastEditTime: 2020-09-21 16:43:23
 */
package wrappers

import (
	"api-gateway/services"
	"context"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/client"
)

//NewCarousel 新建轮播图信息
func NewCarousel(id uint32) *services.CarouselModel {
	return &services.CarouselModel{
		ID:        id,
		ProductID: id,
		ImgPath:   "https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/ac3ef452ba1527d6bf32dc2c3ffc617c.jpg?thumb=1&w=1226&h=460&f=webp&q=90",
	}
}

//DefaultCarousels 降级处理函数，固定生成5个轮播图信息
func DefaultCarousels(rsp interface{}) {
	models := make([]*services.CarouselModel, 0)
	var i uint32
	for i = 0; i < 8; i++ {
		models = append(models, NewCarousel(i))
	}
	result := rsp.(*services.CarouselsListResponse)
	result.CarouselsList = models
}

//DefaultCarouselData 通用降级方法
func DefaultCarouselData(rsp interface{}) {
	switch t := rsp.(type) {
	case *services.CarouselsListResponse:
		DefaultCarousels(rsp)
	case *services.CarouselDetailResponse:
		t.CarouselDetail = NewCarousel(1)
	}
}

//CarouselWrapper go-micro Wrapper轮播图中间件
type CarouselWrapper struct {
	client.Client
}

//Call Wrapper中间件的执行方法
func (this *CarouselWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint()
	config := hystrix.CommandConfig{
		Timeout:                3000,
		RequestVolumeThreshold: 20,   //熔断器请求阈值，默认20，意思是有20个请求才能进行错误百分比计算
		ErrorPercentThreshold:  50,   //错误百分比，当错误超过百分比时，直接进行降级处理，直至熔断器再次 开启，默认50%
		SleepWindow:            5000, //过多长时间，熔断器再次检测是否开启，单位毫秒ms（默认5秒）
	}
	hystrix.ConfigureCommand(cmdName, config)
	return hystrix.Do(cmdName, func() error {
		return this.Client.Call(ctx, req, rsp)
	}, func(e error) error {
		DefaultCarouselData(rsp)
		return nil
	})
}

//NewCarouselWrapper 初始化Wrapper
func NewCarouselWrapper(c client.Client) client.Client {
	return &CarouselWrapper{c}
}
