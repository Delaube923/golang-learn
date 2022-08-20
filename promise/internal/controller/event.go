package controller

import (
	"context"

	v1 "promise/api/v1"
	"promise/internal/service"

	"github.com/gogf/gf/v2/frame/g"
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

	g.Log().Print(ctx, res.List)

	return
}

// AddEvent新增事件
func (c *cEvent) AddEvent(ctx context.Context, req *v1.EventAddReq) (res *v1.EventAddRes, err error) {
	err = service.Event().AddEvent(ctx, req)
	return
}
