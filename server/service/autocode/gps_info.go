package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type GpsInfoService struct {
}

// CreateGpsInfo 创建GpsInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (gpsInfoService *GpsInfoService) CreateGpsInfo(gpsInfo autocode.GpsInfo) (err error) {
	err = global.GVA_DB.Create(&gpsInfo).Error
	return err
}

// DeleteGpsInfo 删除GpsInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (gpsInfoService *GpsInfoService) DeleteGpsInfo(gpsInfo autocode.GpsInfo) (err error) {
	err = global.GVA_DB.Delete(&gpsInfo).Error
	return err
}

// DeleteGpsInfoByIds 批量删除GpsInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (gpsInfoService *GpsInfoService) DeleteGpsInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.GpsInfo{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateGpsInfo 更新GpsInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (gpsInfoService *GpsInfoService) UpdateGpsInfo(gpsInfo autocode.GpsInfo) (err error) {
	err = global.GVA_DB.Save(&gpsInfo).Error
	return err
}

// GetGpsInfo 根据id获取GpsInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (gpsInfoService *GpsInfoService) GetGpsInfo(id uint) (err error, gpsInfo autocode.GpsInfo) {
	err = global.GVA_DB.Where("id = ?", id).First(&gpsInfo).Error
	return
}

// GetGpsInfoInfoList 分页获取GpsInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (gpsInfoService *GpsInfoService) GetGpsInfoInfoList(info autoCodeReq.GpsInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.GpsInfo{})
	var gpsInfos []autocode.GpsInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.GpsSn != "" {
		db = db.Where("`gps_sn` = ?", info.GpsSn)
	}
	if info.BusId != nil {
		db = db.Where("`bus_id` = ?", info.BusId)
	}

	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&gpsInfos).Error
	return err, gpsInfos, total
}

// GetAvailableGpsInfoInfoList 分页获取可用GpsInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (gpsInfoService *GpsInfoService) GetAvailableGpsInfoList(info autoCodeReq.GpsInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.GpsInfo{})
	var gpsInfos []autocode.GpsInfo

	db = db.Where("`bus_id` = ?", -1)

	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&gpsInfos).Error
	return err, gpsInfos, total
}

func (gpsInfoService *GpsInfoService) GetNotAvailableGpsInfoList(info autoCodeReq.GpsInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.GpsInfo{})
	var gpsInfos []autocode.GpsInfo

	db = db.Where("`bus_id` != ?", -1)

	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&gpsInfos).Error
	return err, gpsInfos, total
}
