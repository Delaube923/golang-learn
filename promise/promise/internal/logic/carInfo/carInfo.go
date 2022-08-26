package carinfo

import (
	"context"

	v1 "promise/api/v1"
	"promise/internal/consts"
	"promise/internal/model/do"
	"promise/internal/service"

	"promise/internal/dao"
	"promise/internal/library/liberr"
	"promise/internal/model"
	commonService "promise/internal/service"

	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type sCarInfo struct{}

func init() {
	service.RegisterCarInfo(New())
}

func New() *sCarInfo {
	return &sCarInfo{}
}

type test struct {
	VehicleNumber      string `json:"vehicleNumber" description:"车辆编号"`
	VehicleModel       string `json:"vehicleModel"  description:"车辆型号"`
	SliceName          string `json:"sliceName" description:"切片名称"`
	SliceUrl           string `json:"sliceUrl"      description:"切片存储地址"`
	SliceSize          int    `json:"sliceSize"     description:"切片大小"`
	SliceMd5           string `json:"sliceMd5"      description:"切片md5值"`
	EventId            string `json:"eventId"       description:"事件id"`
	VehicleFrameNumber string `json:"vehicleFrameNumber" description:"车架号"`
	VehicleUsege       string `json:"vehicleUsege" description:"车辆用途"`
	VehicleRegion      string `json:"vehicleRegion" description:"所属地区"`
}

// 从数据库获取车辆信息
func (s *sCarInfo) GetCarInfoListFromDb(ctx context.Context) (value interface{}, err error) {

	err = g.Try(func() {
		var v []*model.CarInfoListItem
		// err = dao.Carinfo.Ctx(ctx).
		// 	Fields(model.CarInfoListItem{}).Order("vehicle_number desc").Scan(&v)
		// fmt.Println(g.DB().Model().Where("vehicle_number", 22))
		// c := g.DB().Model().Where("vehicle_number", 22)
		err := dao.Carinfo.Ctx(ctx).Order(dao.Carinfo.Columns().VehicleNumber).Scan(&v)
		liberr.ErrIsNil(ctx, err, "车辆信息获取失败")
		value = v
		g.Log().Header(false).Print(ctx, v)

	})
	return
}

// 从缓存获取车辆信息
func (s *sCarInfo) GetCarInfoFromCache(ctx context.Context) (list []*model.CarInfoListItem, err error) {
	cache := commonService.Cache()
	iList := cache.GetOrSetFuncLock(ctx, consts.CacheSysCarInfo, s.GetCarInfoListFromDb, 0, consts.CacheSysAuthTag)
	if iList != nil {
		err = gconv.Struct(iList, &list)
		liberr.ErrIsNil(ctx, err)
	}
	return
}

// 添加车辆信息
func (s *sCarInfo) AddCarInfo(ctx context.Context, req *v1.CarInfoAddReq) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		err = g.Try(func() {
			data := do.Carinfo{
				VehicleNumber:      req.VehicleNumber,
				VehicleModle:       req.VehicleModle,
				VehicleFrameNumber: req.VehicleFrameNumber,
				VehicleUsage:       req.VehicleUsage,
				VehicleRegion:      req.VehicleRegion,
			}
			_, e := dao.Carinfo.Ctx(ctx).TX(tx).Insert(data)
			liberr.ErrIsNil(ctx, e, "添加车辆信息失败")

		})
		return err
	})
	if err != nil {
		// 删除相关缓存
		commonService.Cache().Remove(ctx, consts.CacheSysCarInfo)
	}
	return
}
