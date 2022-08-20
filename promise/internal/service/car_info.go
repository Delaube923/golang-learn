// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"context"
	v1 "promise/api/v1"
	"promise/internal/model"
)

type ICarInfo interface {
	GetCarInfoListFromDb(ctx context.Context) (value interface{}, err error)
	GetCarInfoFromCache(ctx context.Context) (list []*model.CarInfoListItem, err error)
	AddCarInfo(ctx context.Context, req *v1.CarInfoAddReq) (err error)
}

var localCarInfo ICarInfo

func CarInfo() ICarInfo {
	if localCarInfo == nil {
		panic("implement not found for interface ICarInfo, forgot register?")
	}
	return localCarInfo
}

func RegisterCarInfo(i ICarInfo) {
	localCarInfo = i
}
