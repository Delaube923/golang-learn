package controller

import (
	"context"
	v1 "promise/api/v1"
	"promise/internal/service"
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
	return
}

// AddCarInfo 添加车辆信息
func (c *cCarInfo) AddCarInfo(ctx context.Context, req *v1.CarInfoAddReq) (res *v1.CarInfoAddRes, err error) {
	err = service.CarInfo().AddCarInfo(ctx, req)
	return
}
