package v1

import (
	"promise/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type CarInfoListReq struct {
	g.Meta        `path:"/carInfo/list" method:"get" tags:"车辆" summary:"车辆列表"`
	VehicleNumber string `json:"vehicle_number"`
	CommonPaginationReq
}
type CarInfoListRes struct {
	g.Meta `mime:"application/json""`
	List   []*model.CarInfoListItem `json:"carList"`
	CommonPaginationRes
}
type CarInfoBase struct {
	VehicleNumber string `json:"vehicleNumber" description:"车辆编号"`
	VehicleModel  string `json:"vehicleModel"  description:"车辆型号"`
	SliceUrl      string `json:"sliceUrl"      description:"切片存储地址"`
	SliceName     string `json:"sliceName" description:"切片名称"`
	SliceSize     int    `json:"sliceSize"     description:"切片大小"`
	SliceMd5      string `json:"sliceMd5"      description:"切片md5值"`
	EventId       string `json:"eventId"       description:"事件id"`
}
type CarInfoAddReq struct {
	g.Meta `path:"/carinfo/add" method:"post" tags:"车辆" summary:"添加车辆信息"`
	CarInfoBase
}
type CarInfoAddRes struct {
}
