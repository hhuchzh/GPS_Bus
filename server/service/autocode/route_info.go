package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type RouteInfoService struct {
}

// CreateRouteInfo 创建RouteInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (routeInfoService *RouteInfoService) CreateRouteInfo(routeInfo autocode.RouteInfo) (err error) {
	err = global.GVA_DB.Omit("LocationInfos").Create(&routeInfo).Error
	return err
}

// DeleteRouteInfo 删除RouteInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (routeInfoService *RouteInfoService) DeleteRouteInfo(routeInfo autocode.RouteInfo) (err error) {
	err = global.GVA_DB.Delete(&routeInfo).Error
	return err
}

// DeleteRouteInfoByIds 批量删除RouteInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (routeInfoService *RouteInfoService) DeleteRouteInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.RouteInfo{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateRouteInfo 更新RouteInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (routeInfoService *RouteInfoService) UpdateRouteInfo(routeInfo autocode.RouteInfo) (err error) {
	err = global.GVA_DB.Omit("LocationInfos").Save(&routeInfo).Error
	return err
}

// GetRouteInfo 根据id获取RouteInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (routeInfoService *RouteInfoService) GetRouteInfo(id uint) (err error, routeInfo autocode.RouteInfo) {
	err = global.GVA_DB.Where("id = ?", id).Preload("LocationInfos").First(&routeInfo).Error
	return
}

// GetRouteInfoInfoList 分页获取RouteInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (routeInfoService *RouteInfoService) GetRouteInfoInfoList(info autoCodeReq.RouteInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.RouteInfo{})
	var routeInfos []autocode.RouteInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.RouteName != "" {
		db = db.Where("`route_name` LIKE ?", "%"+info.RouteName+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("LocationInfos").Find(&routeInfos).Error
	return err, routeInfos, total
}
