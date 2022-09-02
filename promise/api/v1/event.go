package v1

import (
	"promise/internal/model"
	"time"

	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/v2/frame/g"
)

type EventGetListCommonReq struct {
	EventId   int       `json:"eventId" in:"query" dc:"事件id"`
	Type      string    `json:"eventType" in:"query" dc:"事件类型"`
	StartTime time.Time `json:"startTime" in:"query" dc:"事件开始时间"`
	Duration  int       `json:"duration" in:"query" dc:"事件持续时间"`
	CommonPaginationReq
}

type EventGetListCommonRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
	CommonPaginationRes
}

type EventIndexReq struct {
	g.Meta `path:"/event/list" method:"get" tags:"事件" summary:"事件列表"`
	// EventGetListCommonReq
}

type EventIndexRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.EventListItem `json:"data"`
	// CommonPaginationRes
}

// 事件列表结构体
type EventBase struct {
	Id               int         `json:"eventId"          description:"事件id"`
	VehicleNumber    string      `json:"vehicleNumber"   description:"车辆编号"`
	VehicleModel     string      `json:"vehicleModel"    description:"车辆类型"`
	EventTime        *gtime.Time `json:"eventTime"        description:"事件发生日期/时间"`
	TriggerType      string      `json:"triggerType"      description:"事件触发方式"`
	EventType        string      `json:"eventType"        description:"事件类型"`
	EventDescription string      `json:"eventDescription" description:"事件描述"`
	StartTime        *gtime.Time `json:"startTime"        description:"切片数据开始时间"`
	Duration         int         `json:"duration"         description:"切片数据的持续时间(s)"`
	SliceUrl         string      `json:"sliceUrl"         description:"切片存储地址"`
	SliceName        string      `json:"sliceName"        description:"切片名称"`
	SliceSize        int         `json:"sliceSize"        description:"切片大小"`
	SliceMd5         string      `json:"sliceMd5"         description:"切片md5值"`
}

type EventAddReq struct {
	g.Meta           `path:"/event/add" method:"post" tags:"事件" summary:"添加事件"`
	EventId          int         `json:"eventId" description:"事件id" gorm:"-;primary_key;AUTO_INCREMENT"`
	VehicleNumber    string      `json:"vehicleNumber"   description:"车辆编号"`
	VehicleModel     string      `json:"vehicleModel"   description:"车辆类型"`
	EventTime        *gtime.Time `json:"eventTime"        description:"事件发生日期/时间"`
	TriggerType      string      `json:"triggerType"      description:"事件触发方式"`
	EventType        string      `json:"eventType"        description:"事件类型"`
	EventDescription string      `json:"eventDescription" description:"事件描述"`
	StartTime        *gtime.Time `json:"startTime"        description:"切片数据开始时间"`
	Duration         int         `json:"duration"         description:"切片数据的持续时间(s)"`
	SliceUrl         string      `json:"sliceUrl"         description:"切片存储地址"`
	SliceName        string      `json:"sliceName"        description:"切片名称"`
	SliceSize        int         `json:"sliceSize"        description:"切片大小"`
	SliceMd5         string      `json:"sliceMd5"         description:"切片md5值"`
}
type EventAddRes struct {
}
type EventSearchReq struct {
	g.Meta        `path:"/event/search" method:"get" tags:"事件" summary:"条件查询事件"`
	VehicleNumber string      `json:"vehicleNumber"   description:"车辆编号"`
	EventTime     *gtime.Time `json:"eventTime"        description:"事件发生日期/时间"`
	DateRange     []string    `json:"dataRange" description:"日期范围"`
}
type EventSearchRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.EventListItem `json:"data"`
}
