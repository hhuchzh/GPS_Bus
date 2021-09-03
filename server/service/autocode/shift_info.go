package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type ShiftInfoService struct {
}

// CreateShiftInfo 创建ShiftInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (shiftInfoService *ShiftInfoService) CreateShiftInfo(shiftInfo autocode.ShiftInfo) (err error) {
	err = global.GVA_DB.Create(&shiftInfo).Error
	return err
}

// DeleteShiftInfo 删除ShiftInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (shiftInfoService *ShiftInfoService)DeleteShiftInfo(shiftInfo autocode.ShiftInfo) (err error) {
	err = global.GVA_DB.Delete(&shiftInfo).Error
	return err
}

// DeleteShiftInfoByIds 批量删除ShiftInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (shiftInfoService *ShiftInfoService)DeleteShiftInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.ShiftInfo{},"id in ?",ids.Ids).Error
	return err
}

// UpdateShiftInfo 更新ShiftInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (shiftInfoService *ShiftInfoService)UpdateShiftInfo(shiftInfo autocode.ShiftInfo) (err error) {
	err = global.GVA_DB.Save(&shiftInfo).Error
	return err
}

// GetShiftInfo 根据id获取ShiftInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (shiftInfoService *ShiftInfoService)GetShiftInfo(id uint) (err error, shiftInfo autocode.ShiftInfo) {
	err = global.GVA_DB.Where("id = ?", id).First(&shiftInfo).Error
	return
}

// GetShiftInfoInfoList 分页获取ShiftInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (shiftInfoService *ShiftInfoService)GetShiftInfoInfoList(info autoCodeReq.ShiftInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.ShiftInfo{})
    var shiftInfos []autocode.ShiftInfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.BusId != nil {
        db = db.Where("`bus_id` = ?",info.BusId)
    }
    if info.ClassesId != nil {
        db = db.Where("`classes_id` = ?",info.ClassesId)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&shiftInfos).Error
	return err, shiftInfos, total
}
