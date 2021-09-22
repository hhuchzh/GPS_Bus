package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BusInfoService struct {
}

// CreateBusInfo 创建BusInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (busInfoService *BusInfoService) CreateBusInfo(busInfo autocode.BusInfo) (err error) {
	err = global.GVA_DB.Omit("GpsInfos").Create(&busInfo).Error
	return err
}

// DeleteBusInfo 删除BusInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (busInfoService *BusInfoService) DeleteBusInfo(busInfo autocode.BusInfo) (err error) {
	err = global.GVA_DB.Delete(&busInfo).Error
	return err
}

// DeleteBusInfoByIds 批量删除BusInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (busInfoService *BusInfoService) DeleteBusInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.BusInfo{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateBusInfo 更新BusInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (busInfoService *BusInfoService) UpdateBusInfo(busInfo autocode.BusInfo) (err error) {
	err = global.GVA_DB.Omit("GpsInfos").Save(&busInfo).Error
	return err
}

// GetBusInfo 根据id获取BusInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (busInfoService *BusInfoService) GetBusInfo(id uint) (err error, busInfo autocode.BusInfo) {
	err = global.GVA_DB.Where("id = ?", id).Preload("GpsInfos").First(&busInfo).Error
	return
}

// GetBusInfoInfoList 分页获取BusInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (busInfoService *BusInfoService) GetBusInfoInfoList(info autoCodeReq.BusInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.BusInfo{})
	var busInfos []autocode.BusInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.BusUserPhone != "" {
		db = db.Where("`bus_user_phone` = ?", info.BusUserPhone)
	}
	if info.BusUserName != "" {
		db = db.Where("`bus_user_name` = ?", info.BusUserName)
	}
	if info.BusPlate != "" {
		db = db.Where("`bus_plate` = ?", info.BusPlate)
	}
	if info.CompanyId != nil {
		db = db.Where("`company_id` = ?", info.CompanyId)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("GpsInfos").Find(&busInfos).Error
	return err, busInfos, total
}
