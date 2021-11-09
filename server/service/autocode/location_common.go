package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type LocationCommonService struct {
}

// CreateLocationCommon 创建LocationCommon记录
// Author [piexlmax](https://github.com/piexlmax)
func (locationCommonService *LocationCommonService) CreateLocationCommon(locationCommon autocode.LocationCommon) (err error) {
	err = global.GVA_DB.Create(&locationCommon).Error
	return err
}

// DeleteLocationCommon 删除LocationCommon记录
// Author [piexlmax](https://github.com/piexlmax)
func (locationCommonService *LocationCommonService)DeleteLocationCommon(locationCommon autocode.LocationCommon) (err error) {
	err = global.GVA_DB.Delete(&locationCommon).Error
	return err
}

// DeleteLocationCommonByIds 批量删除LocationCommon记录
// Author [piexlmax](https://github.com/piexlmax)
func (locationCommonService *LocationCommonService)DeleteLocationCommonByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.LocationCommon{},"id in ?",ids.Ids).Error
	return err
}

// UpdateLocationCommon 更新LocationCommon记录
// Author [piexlmax](https://github.com/piexlmax)
func (locationCommonService *LocationCommonService)UpdateLocationCommon(locationCommon autocode.LocationCommon) (err error) {
	err = global.GVA_DB.Save(&locationCommon).Error
	return err
}

// GetLocationCommon 根据id获取LocationCommon记录
// Author [piexlmax](https://github.com/piexlmax)
func (locationCommonService *LocationCommonService)GetLocationCommon(id uint) (err error, locationCommon autocode.LocationCommon) {
	err = global.GVA_DB.Where("id = ?", id).First(&locationCommon).Error
	return
}

// GetLocationCommonInfoList 分页获取LocationCommon记录
// Author [piexlmax](https://github.com/piexlmax)
func (locationCommonService *LocationCommonService)GetLocationCommonInfoList(info autoCodeReq.LocationCommonSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.LocationCommon{})
    var locationCommons []autocode.LocationCommon
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.LocationName != "" {
        db = db.Where("`location_name` LIKE ?","%"+ info.LocationName+"%")
    }
    if info.Longtitude != "" {
        db = db.Where("`longtitude` = ?",info.Longtitude)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&locationCommons).Error
	return err, locationCommons, total
}
