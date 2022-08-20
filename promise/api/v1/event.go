package v1

import (
	"promise/internal/model"
	"time"

	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/v2/frame/g"
)

type EventGetListCommonReq struct {
	EventId   string    `json:"eventId" in:"query" dc:"事件id"`
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
	EventGetListCommonReq
}

type EventIndexRes struct {
	g.Meta `mime:"application/json""`
	List   []*model.EventListItem `json:"eventlist"`
	CommonPaginationRes
}
type EventBaseReq struct {
	g.Meta           `path:"/event/addByChannel" method:"post" `
	Id               string      `json:"eventId"               description:"事件id"`
	VehicleNumber    string      `json:"vehicle_number"   description:"车辆编号"`
	VehicleModel     string      `json:"vehicle_model"   description:"车辆类型"`
	EventTime        *gtime.Time `json:"eventTime"        description:"事件发生日期/时间"`
	TriggerType      string      `json:"triggerType"      description:"事件触发方式"`
	EventType        string      `json:"eventType"        description:"事件类型"`
	EventDescription string      `json:"eventDescription" description:"事件描述"`
	StartTime        *gtime.Time `json:"startTime"        description:"切片数据开始时间"`
	Duration         int         `json:"duration"         description:"切片数据的持续时间(s)"`
}

// 事件列表结构体
type EventBase struct {
	Id               string      `json:"eventId"               description:"事件id"`
	VehicleNumber    string      `json:"vehicle_number"   description:"车辆编号"`
	VehicleModel     string      `json:"vehicle_model"   description:"车辆类型"`
	EventTime        *gtime.Time `json:"eventTime"        description:"事件发生日期/时间"`
	TriggerType      string      `json:"triggerType"      description:"事件触发方式"`
	EventType        string      `json:"eventType"        description:"事件类型"`
	EventDescription string      `json:"eventDescription" description:"事件描述"`
	StartTime        *gtime.Time `json:"startTime"        description:"切片数据开始时间"`
	Duration         int         `json:"duration"         description:"切片数据的持续时间(s)"`
}

type EventAddReq struct {
	g.Meta           `path:"/event/add" method:"post" tags:"事件" summary:"添加事件"`
	Id               string      `json:"eventId" description:"事件id" gorm:"-;primary_key;AUTO_INCREMENT"`
	VehicleNumber    string      `json:"vehicle_number"   description:"车辆编号"`
	VehicleModel     string      `json:"vehicle_model"   description:"车辆类型"`
	EventTime        *gtime.Time `json:"eventTime"        description:"事件发生日期/时间"`
	TriggerType      string      `json:"triggerType"      description:"事件触发方式"`
	EventType        string      `json:"eventType"        description:"事件类型"`
	EventDescription string      `json:"eventDescription" description:"事件描述"`
	StartTime        *gtime.Time `json:"startTime"        description:"切片数据开始时间"`
	Duration         int         `json:"duration"         description:"切片数据的持续时间(s)"`
}
type EventAddRes struct {
}
