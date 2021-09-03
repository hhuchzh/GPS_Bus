package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type CheckinInfoService struct {
}

// CreateCheckinInfo 创建CheckinInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkinInfoService *CheckinInfoService) CreateCheckinInfo(checkinInfo autocode.CheckinInfo) (err error) {
	err = global.GVA_DB.Create(&checkinInfo).Error
	return err
}

// DeleteCheckinInfo 删除CheckinInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkinInfoService *CheckinInfoService)DeleteCheckinInfo(checkinInfo autocode.CheckinInfo) (err error) {
	err = global.GVA_DB.Delete(&checkinInfo).Error
	return err
}

// DeleteCheckinInfoByIds 批量删除CheckinInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkinInfoService *CheckinInfoService)DeleteCheckinInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.CheckinInfo{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCheckinInfo 更新CheckinInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkinInfoService *CheckinInfoService)UpdateCheckinInfo(checkinInfo autocode.CheckinInfo) (err error) {
	err = global.GVA_DB.Save(&checkinInfo).Error
	return err
}

// GetCheckinInfo 根据id获取CheckinInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkinInfoService *CheckinInfoService)GetCheckinInfo(id uint) (err error, checkinInfo autocode.CheckinInfo) {
	err = global.GVA_DB.Where("id = ?", id).First(&checkinInfo).Error
	return
}

// GetCheckinInfoInfoList 分页获取CheckinInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkinInfoService *CheckinInfoService)GetCheckinInfoInfoList(info autoCodeReq.CheckinInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.CheckinInfo{})
    var checkinInfos []autocode.CheckinInfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.LocationId != nil {
        db = db.Where("`location_id` = ?",info.LocationId)
    }
    if info.ShiftId != nil {
        db = db.Where("`shift_id` = ?",info.ShiftId)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&checkinInfos).Error
	return err, checkinInfos, total
}
