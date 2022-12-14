// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EventmiddleDao is the data access object for table eventmiddle.
type EventmiddleDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns EventmiddleColumns // columns contains all the column names of Table for convenient usage.
}

// EventmiddleColumns defines and stores column names for table eventmiddle.
type EventmiddleColumns struct {
	EventId          string // 事件id
	EventTime        string // 事件发生日期/时间
	EventType        string // 事件类型
	EventDescription string // 事件描述
	StartTime        string // 数据开始时间
	Duration         string // 切片数据的持续时间(s)
	TriggerType      string // 事件触发方式
	VehicleNumber    string // 车辆编号
	VehicleModel     string // 车辆型号
	SliceUrl         string // 切片存储地址
	SliceName        string // 切片名称
	SliceSize        string // 切片大小
	SliceMd5         string // 切片md5值
}

//  eventmiddleColumns holds the columns for table eventmiddle.
var eventmiddleColumns = EventmiddleColumns{
	EventId:          "eventId",
	EventTime:        "event_time",
	EventType:        "event_type",
	EventDescription: "event_description",
	StartTime:        "start_time",
	Duration:         "duration",
	TriggerType:      "trigger_type",
	VehicleNumber:    "vehicle_number",
	VehicleModel:     "vehicle_model",
	SliceUrl:         "slice_url",
	SliceName:        "slice_name",
	SliceSize:        "slice_size",
	SliceMd5:         "slice_md5",
}

// NewEventmiddleDao creates and returns a new DAO object for table data access.
func NewEventmiddleDao() *EventmiddleDao {
	return &EventmiddleDao{
		group:   "default",
		table:   "eventmiddle",
		columns: eventmiddleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EventmiddleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EventmiddleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EventmiddleDao) Columns() EventmiddleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EventmiddleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EventmiddleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EventmiddleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
