package event

import (
	"context"
	"fmt"

	v1 "promise/api/v1"
	"promise/internal/consts"
	"promise/internal/dao"
	"promise/internal/library/liberr"
	"promise/internal/model"
	"promise/internal/model/do"
	"promise/internal/service"
	commonService "promise/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sEvent struct{}

func init() {
	service.RegisterEvent(New())
}

func New() *sEvent {
	return &sEvent{}
}

// AddEvent 新建事件
func (s *sEvent) AddEvent(ctx context.Context, req *v1.EventAddReq) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		err = g.Try(func() {
			//事件数据
			data := do.Eventsmall{
				EventId:          req.Id,
				VehicleNumber:    req.VehicleNumber,
				TriggerType:      req.TriggerType,
				EventType:        req.EventType,
				EventDescription: req.EventDescription,
				Duration:         req.Duration,
			}
			EventId, e := dao.Eventsmall.Ctx(ctx).TX(tx).InsertAndGetId(data)
			liberr.ErrIsNil(ctx, e, "添加事件失败")
			print(EventId)
		})
		return err
	})
	if err == nil {
		// 删除相关缓存
		commonService.Cache().Remove(ctx, consts.CacheSysEvent)
	}
	return
}

// AddEventByChannel 从管道读取数据并新建事件
func (s *sEvent) AddEventByChannel(ctx context.Context, req *v1.EventBaseReq) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		err = g.Try(func() {
			data := do.Eventsmall{
				VehicleNumber: req.VehicleNumber,
				// EventTime: req.EventTime,
				TriggerType: req.TriggerType,
				EventType:   req.EventType,
				Duration:    req.Duration,
			}
			_, e := dao.Eventsmall.Ctx(ctx).TX(tx).Insert(data)
			liberr.ErrIsNil(ctx, e, "添加事件失败")

		})
		return err
	})
	if err == nil {
		commonService.Cache().Remove(ctx, consts.CacheSysEvent)
	}
	return
}

// 从数据库获取所有事件（周表）
func (s *sEvent) GetWeekEventListFromDb(ctx context.Context) (value interface{}, err error) {
	err = g.Try(func() {
		var v []*model.EventListItem
		//从数据库获取,事件时间降序
		err = dao.Eventsmall.Ctx(ctx).
			Fields(model.EventListItem{}).Order("event_time desc").Scan(&v)
		liberr.ErrIsNil(ctx, err, "最近一周数据获取失败")
		value = v
		fmt.Print("111", value)
	})
	return
}

// 从缓存获取最近事件
func (s *sEvent) GetWeekEventListFromCache(ctx context.Context) (list []*model.EventListItem, err error) {
	cache := commonService.Cache()
	iList := cache.GetOrSetFuncLock(ctx, consts.CacheSysEvent, s.GetWeekEventListFromDb, 0, consts.CacheSysAuthTag)
	if iList != nil {
		err = gconv.Struct(iList, &list)
		liberr.ErrIsNil(ctx, err)
	}
	return
}
