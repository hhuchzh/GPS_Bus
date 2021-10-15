package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ClassesInfoService struct {
}

// CreateClassesInfo 创建ClassesInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (classesInfoService *ClassesInfoService) CreateClassesInfo(classesInfo autocode.ClassesInfo) (err error) {
	err = global.GVA_DB.Omit("Bus").Create(&classesInfo).Error
	return err
}

// DeleteClassesInfo 删除ClassesInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (classesInfoService *ClassesInfoService) DeleteClassesInfo(classesInfo autocode.ClassesInfo) (err error) {
	err = global.GVA_DB.Delete(&classesInfo).Error
	return err
}

// DeleteClassesInfoByIds 批量删除ClassesInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (classesInfoService *ClassesInfoService) DeleteClassesInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.ClassesInfo{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateClassesInfo 更新ClassesInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (classesInfoService *ClassesInfoService) UpdateClassesInfo(classesInfo autocode.ClassesInfo) (err error) {
	err = global.GVA_DB.Omit("Bus").Save(&classesInfo).Error
	return err
}

// GetClassesInfo 根据id获取ClassesInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (classesInfoService *ClassesInfoService) GetClassesInfo(id uint) (err error, classesInfo autocode.ClassesInfo) {
	err = global.GVA_DB.Where("id = ?", id).Preload("Bus").First(&classesInfo).Error
	return
}

// GetClassesInfoInfoList 分页获取ClassesInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (classesInfoService *ClassesInfoService) GetClassesInfoInfoList(info autoCodeReq.ClassesInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.ClassesInfo{})
	var classesInfos []autocode.ClassesInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.RouteId != nil {
		db = db.Where("`route_id` = ?", info.RouteId)
	}

	if info.BusId != nil {
		db = db.Where("`bus_id` = ?", info.BusId)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Bus").Find(&classesInfos).Error
	return err, classesInfos, total
}
