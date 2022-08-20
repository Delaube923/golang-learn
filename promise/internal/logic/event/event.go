package event

import (
	"context"

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
	"github.com/gogf/gf/v2/os/gtime"
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
				EventId:          req.EventId,
				EventTime:        gtime.NewFromTime(req.EventTime.Time),
				EventType:        req.EventType,
				EventDescription: req.EventDescription,
				StartTime:        gtime.NewFromTime(req.StartTime.Time),
				Duration:         req.Duration,
				TriggerType:      req.TriggerType,
				VehicleNumber:    req.VehicleNumber,
				VehicleModel:     req.VehicleModel,
			}
			// data := do.Eventsmall{
			// 	EventId:          8989,
			// 	VehicleNumber:    0505,
			// 	TriggerType:      "人为触发",
			// 	EventType:        "超车",
			// 	EventDescription: 0,
			// 	Duration:         10,
			// }
			EventId, e := dao.Eventsmall.Ctx(ctx).TX(tx).InsertAndGetId(data)
			liberr.ErrIsNil(ctx, e, "添加事件失败")
			print(EventId)
			// _, e := dao.Eventsmall.Ctx(ctx).TX(tx).Data(g.Map{"eventId": 6666}).Insert()
			g.Log().Print(ctx, e)
		})
		return err
	})
	if err == nil {
		// 删除相关缓存
		commonService.Cache().Remove(ctx, consts.CacheSysEvent)
	}
	return
}

// 从数据库获取所有事件（周表）
func (s *sEvent) GetWeekEventListFromDb(ctx context.Context) (value interface{}, err error) {
	err = g.Try(func() {
		var v []*model.EventListItem
		//从数据库获取,事件时间降序
		// err = dao.Eventsmall.Ctx(ctx).
		// 	Fields(model.EventListItem{}).Order("event_time desc").Scan(&v)
		err := dao.Eventsmall.Ctx(ctx).Order(dao.Eventsmall.Columns().EventTime).Scan(&v)

		liberr.ErrIsNil(ctx, err, "最近一周数据获取失败")
		value = v
		g.Log().Header(false).Print(ctx, value)
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
	// g.Log().Print(ctx, iList)
	// fmt.Println(11111111111)
	return
}
