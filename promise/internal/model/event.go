package model

import (
	"time"

	"github.com/gogf/gf/v2/os/gtime"
)

// EventGetListInput获取事件信息
type EventGetListInput struct {
	EventId   string    //事件id
	Type      string    //事件类型
	StartTime time.Time //事件开始时间
	Duration  int       //事件持续时间
	Page      int       //分页号码
	Size      int       //分页数量，最大50
}

// 查询列表结果
type EventGetListOutput struct {
	List  []EventListOutputItem `json:"data" description:"列表"`
	Page  int                   `json:"page" description:"分页码"`
	Size  int                   `json:"size" description:"分页数量"`
	Total int                   `json:"total" description:"数据总数"`
}

// 查询结果项
type EventListOutputItem struct {
	Event *EventListItem `json:"event"`
}

// 列表信息
type EventListItem struct {
	EventId          int         `json:"eventId" description:"事件id"`
	EventTime        *gtime.Time `json:"eventTime" description:"事件发生日期/时间"`
	EventType        string      `json:"eventType" description:"事件类型"`
	EventDescription string      `json:"eventDescription" description:"事件描述"`
	TriggerType      string      `json:"TriggerType" description:"触发类型"`
	VehicleNumber    string      `json:"vehicleNumber"   description:"车辆编号"`
	VehicleModel     string      `json:"vehicleModel"   description:"车辆类型"`
	SliceUrl         string      `json:"sliceUrl"         description:"切片存储地址"`
	SliceName        string      `json:"sliceName"        description:"切片名称"`
	SliceSize        int         `json:"sliceSize"        description:"切片大小"`
	SliceMd5         string      `json:"sliceMd5"         description:"切片md5值"`
	StartTime        *gtime.Time `json:"startTime"  description:"切片数据开始时间"`
	Duration         int         `json:"duration" description:"切片数据的持续时间(s)"`
}

// 添加事件
type EventCreateInput struct {
	EventListItem
}
