package model

//CarInfoGetListInput  获取车辆信息
type CarInfoGetListInput struct {
	CarId string //要查询的车辆id
	Type  string //要查询的车辆型号
	Area  string //要查询的车辆地区
	Use   string //要查询的车辆用途
	Page  int    //分页号码
	Size  int    //分页数量，最大50

}

//CarInfoGetListOutput 查询列表结果
type CarInfoGetListOutput struct {
	List  []CarInfoListOutputItem `json:"carlist" description:"列表"`
	Page  int                     `json:"page" description:"分页码"`
	Size  int                     `json:"size" description:"分页数量"`
	Total int                     `json:"total" description:"数据总数"`
}

//查询列表结果项
type CarInfoListOutputItem struct {
	Car *CarInfoListItem `json:"car"`
}

//列表信息
type CarInfoListItem struct {
	VehicleNumber string `json:"vehicleNumber" description:"车辆编号"`
	VehicleModel  string `json:"vehicleModel"  description:"车辆型号"`
	SliceName     string `json:"sliceName" description:"切片名称"`
	SliceUrl      string `json:"sliceUrl"      description:"切片存储地址"`
	SliceSize     int    `json:"sliceSize"     description:"切片大小"`
	SliceMd5      string `json:"sliceMd5"      description:"切片md5值"`
	EventId       string `json:"eventId"       description:"事件id"`
}
