package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type LocationInfoService struct {
}

// CreateLocationInfo 创建LocationInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (locationInfoService *LocationInfoService) CreateLocationInfo(locationInfo autocode.LocationInfo) (err error) {
	err = global.GVA_DB.Create(&locationInfo).Error
	return err
}

// DeleteLocationInfo 删除LocationInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (locationInfoService *LocationInfoService)DeleteLocationInfo(locationInfo autocode.LocationInfo) (err error) {
	err = global.GVA_DB.Delete(&locationInfo).Error
	return err
}

// DeleteLocationInfoByIds 批量删除LocationInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (locationInfoService *LocationInfoService)DeleteLocationInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.LocationInfo{},"id in ?",ids.Ids).Error
	return err
}

// UpdateLocationInfo 更新LocationInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (locationInfoService *LocationInfoService)UpdateLocationInfo(locationInfo autocode.LocationInfo) (err error) {
	err = global.GVA_DB.Save(&locationInfo).Error
	return err
}

// GetLocationInfo 根据id获取LocationInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (locationInfoService *LocationInfoService)GetLocationInfo(id uint) (err error, locationInfo autocode.LocationInfo) {
	err = global.GVA_DB.Where("id = ?", id).First(&locationInfo).Error
	return
}

// GetLocationInfoInfoList 分页获取LocationInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (locationInfoService *LocationInfoService)GetLocationInfoInfoList(info autoCodeReq.LocationInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.LocationInfo{})
    var locationInfos []autocode.LocationInfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.LocationName != "" {
        db = db.Where("`location_name` LIKE ?","%"+ info.LocationName+"%")
    }
    if info.RouteId != nil {
        db = db.Where("`route_id` = ?",info.RouteId)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&locationInfos).Error
	return err, locationInfos, total
}
