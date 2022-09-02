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
	VehicleNumber      string `json:"vehicleNumber" description:"车辆编号"`
	VehicleModel       string `json:"vehicleModel"  description:"车辆型号"`
	VehicleFrameNumber string `json:"vehicleFrameNumber" description:"车架号"`
	VehicleUsage       string `json:"vehicleUsage" description:"车辆用途"`
	VehicleRegion      string `json:"vehicleRegion" description:"所属地区"`
	Version            string `json:"version"            description:"大版本号"`
	Status             string `json:"status"             description:"车辆状态"`
}
type CarInfoAddReq struct {
	g.Meta `path:"/carinfo/add" method:"post" tags:"车辆" summary:"添加车辆信息"`
	CarInfoBase
}
type CarInfoAddRes struct {
}
