package controller

import (
	"context"
	v1 "promise/api/v1"
	"promise/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	CarInfo = cCarInfo{}
)

type cCarInfo struct{}

// Get获取车辆信息
func (c *cCarInfo) Get(ctx context.Context, req *v1.CarInfoListReq) (res *v1.CarInfoListRes, err error) {
	// fmt.Println("get....")
	res = new(v1.CarInfoListRes)
	res.List, err = service.CarInfo().GetCarInfoFromCache(ctx)
	if err != nil {
		return nil, err
	}
	g.RequestFromCtx(ctx).Response.WriteJson(res)

	return
}

// AddCarInfo 添加车辆信息
func (c *cCarInfo) AddCarInfo(ctx context.Context, req *v1.CarInfoAddReq) (res *v1.CarInfoAddRes, err error) {
	err = service.CarInfo().AddCarInfo(ctx, req)
	g.RequestFromCtx(ctx).Response.WriteJson(res)
	return
}
