package controller

import (
	"context"

	v1 "promise/api/v1"
	"promise/internal/service"
	"promise/internal/socket"
	"promise/utility/utils"

	"github.com/gogf/gf/os/gtime"
)

// 事件管理
var (
	Event = cEvent{}
)

type cEvent struct{}

// Get获取事件
func (c *cEvent) Get(ctx context.Context, req *v1.EventIndexReq) (res *v1.EventIndexRes, err error) {
	res = new(v1.EventIndexRes)

	res.List, err = service.Event().GetWeekEventListFromCache(ctx)
	if err != nil {
		return nil, err
	}
	print(res.List)

	return
}

// AddEvent新增事件
func (c *cEvent) AddEvent(ctx context.Context, req *v1.EventAddReq) (res *v1.EventAddRes, err error) {
	err = service.Event().AddEvent(ctx, req)
	return
}

// 从channel读取数据 写入mysql
func (c *cEvent) AddEventByChannel(ctx context.Context, req *v1.EventBaseReq) (res *v1.EventAddRes, err error) {
	var x = <-socket.VechicleEventChannel
	eventtime := utils.Int2Time(utils.Interface2Int64(x.EventTime))
	starttime := utils.Int2Time(utils.Interface2Int64(x.StartTime))

	req.VehicleNumber = utils.Interface2String(x.VehicleNumber)
	req.EventTime = gtime.NewFromTime(eventtime)
	req.TriggerType = utils.Interface2String(x.TriggerType)
	req.EventType = utils.Interface2String(x.EventType)
	req.StartTime = gtime.NewFromTime(starttime)
	req.Duration = utils.Interface2Int32(x.Duration)
	err = service.Event().AddEventByChannel(ctx, req)
	return

}
