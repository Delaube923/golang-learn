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
	VehicleNumber      string `json:"vehicleNumber" description:"车辆编号"`
	VehicleModel       string `json:"vehicleModel"  description:"车辆型号"`
	VehicleFrameNumber string `json:"vehicleFrameNumber" description:"车架号"`
	VehicleUsage       string `json:"vehicleUsage" description:"车辆用途"`
	VehicleRegion      string `json:"vehicleRegion" description:"所属地区"`
	Version            string `json:"version"            description:"大版本号"`
	Status             string `json:"status"             description:"车辆状态"`
}
