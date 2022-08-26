// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Eventmiddle is the golang structure of table eventmiddle for DAO operations like Where/Data.
type Eventmiddle struct {
	g.Meta           `orm:"table:eventmiddle, do:true"`
	EventId          interface{} // 事件id
	EventTime        *gtime.Time // 事件发生日期/时间
	EventType        interface{} // 事件类型
	EventDescription interface{} // 事件描述
	StartTime        *gtime.Time // 数据开始时间
	Duration         interface{} // 切片数据的持续时间(s)
	TriggerType      interface{} // 事件触发方式
	VehicleNumber    interface{} // 车辆编号
	VehicleModel     interface{} // 车辆型号
	SliceUrl         interface{} // 切片存储地址
	SliceName        interface{} // 切片名称
	SliceSize        interface{} // 切片大小
	SliceMd5         interface{} // 切片md5值
}