package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ArrivalInfoService struct {
}

// CreateArrivalInfo 创建ArrivalInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (arrivalInfoService *ArrivalInfoService) CreateArrivalInfo(arrivalInfo autocode.ArrivalInfo) (err error) {
	err = global.GVA_DB.Omit("Location").Create(&arrivalInfo).Error
	return err
}

// DeleteArrivalInfo 删除ArrivalInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (arrivalInfoService *ArrivalInfoService) DeleteArrivalInfo(arrivalInfo autocode.ArrivalInfo) (err error) {
	err = global.GVA_DB.Delete(&arrivalInfo).Error
	return err
}

// DeleteArrivalInfoByIds 批量删除ArrivalInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (arrivalInfoService *ArrivalInfoService) DeleteArrivalInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.ArrivalInfo{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateArrivalInfo 更新ArrivalInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (arrivalInfoService *ArrivalInfoService) UpdateArrivalInfo(arrivalInfo autocode.ArrivalInfo) (err error) {
	err = global.GVA_DB.Omit("Location").Save(&arrivalInfo).Error
	return err
}

// GetArrivalInfo 根据id获取ArrivalInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (arrivalInfoService *ArrivalInfoService)GetArrivalInfo(classId *uint, locationId *uint) (err error, arrivalInfo autocode.ArrivalInfo) {
	err = global.GVA_DB.Where("classes_id = ? AND location_id = ?", classId, locationId).First(&arrivalInfo).Error
	return
}

// GetArrivalInfoInfoList 分页获取ArrivalInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (arrivalInfoService *ArrivalInfoService) GetArrivalInfoInfoList(info autoCodeReq.ArrivalInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.ArrivalInfo{})
	var arrivalInfos []autocode.ArrivalInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.LocationId != nil {
		db = db.Where("`location_id` = ?", info.LocationId)
	}
	if info.ClassesId != nil {
		db = db.Where("`classes_id` = ?", info.ClassesId)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Location").Find(&arrivalInfos).Error
	return err, arrivalInfos, total
}
