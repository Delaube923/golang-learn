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
	g.RequestFromCtx(ctx).Response.WriteJson(res)
	return
}

// AddEvent新增事件
func (c *cEvent) AddEvent(ctx context.Context, req *v1.EventAddReq) (res *v1.EventAddRes, err error) {
	err = service.Event().AddEvent(ctx, req)
	g.RequestFromCtx(ctx).Response.WriteJson(res)
	return
}

// // SelectEvent条件查询事件
func (c *cEvent) SelectEvent(ctx context.Context, req *v1.EventSearchReq) (res *v1.EventSearchRes, err error) {
	res = new(v1.EventSearchRes)
	//顺序查询（周表->月表->总表）
	res.List, err = service.Event().GetWeekEventListSearch(ctx, req)
	if res.List == nil {
		res.List, err = service.Event().GetMonthEventListSearch(ctx, req)
		if res.List == nil {
			res.List, err = service.Event().GetGeneralEventListSearch(ctx, req)
		}

	}
	if err != nil {
		return
	}
	g.RequestFromCtx(ctx).Response.WriteJson(res)

	return
}
