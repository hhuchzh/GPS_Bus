// 自动生成模板LocationInfo
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

// LocationInfo 结构体
// 如果含有time.Time 请自行import time包
type LocationInfo struct {
	global.GVA_MODEL
	CreateUserId *int   `json:"createUserId" form:"createUserId" gorm:"column:create_user_id;comment:创建人ID;type:bigint"`
	UpdateUserId *int   `json:"updateUserId" form:"updateUserId" gorm:"column:update_user_id;comment:修改人ID;type:bigint"`
	LocationName string `json:"locationName" form:"locationName" gorm:"column:location_name;comment:站点名;type:varchar(100);"`
	RouteId      *int   `json:"routeId" form:"routeId" gorm:"column:route_id;comment:;type:bigint"`
	Longtitude   string `json:"longtitude" form:"longtitude" gorm:"column:longtitude;comment:经度;type:varchar(50);"`
	Latitude     string `json:"latitude" form:"latitude" gorm:"column:latitude;comment:纬度;type:varchar(50);"`
	OrderNo      *int   `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:站点在路线中的位置，第几站;type:int"`
}

// TableName LocationInfo 表名
func (LocationInfo) TableName() string {
	return "location_info"
}

func (location *LocationInfo) AfterDelete(tx *gorm.DB) (err error) {
	err = tx.Where("location_id = ?", location.ID).Delete(&ArrivalInfo{}).Error
	return err
}
