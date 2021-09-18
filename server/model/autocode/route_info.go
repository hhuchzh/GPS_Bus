// 自动生成模板RouteInfo
package autocode

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

// RouteInfo 结构体
// 如果含有time.Time 请自行import time包
type RouteInfo struct {
	global.GVA_MODEL
	CreateUserId  *int           `json:"createUserId" form:"createUserId" gorm:"column:create_user_id;comment:创建人ID;type:bigint"`
	UpdateUserId  *int           `json:"updateUserId" form:"updateUserId" gorm:"column:update_user_id;comment:修改人ID;type:bigint"`
	RouteName     string         `json:"routeName" form:"routeName" gorm:"column:route_name;comment:路线名;type:varchar(100);"`
	StartAddress  string         `json:"startAddress" form:"startAddress" gorm:"column:start_address;comment:起始地点;type:varchar(100);"`
	EndAddress    string         `json:"endAddress" form:"endAddress" gorm:"column:end_address;comment:结束地点;type:varchar(100);"`
	AboutTime     *time.Time     `json:"aboutTime" form:"aboutTime" gorm:"column:about_time;comment:大约时间;type:time"`
	AboutDistance *int           `json:"aboutDistance" form:"aboutDistance" gorm:"column:about_distance;comment:大约距离;type:int"`
	LocationInfos []LocationInfo `gorm:"foreignKey:RouteId"`
}

// TableName RouteInfo 表名
func (RouteInfo) TableName() string {
	return "route_info"
}

func (classes *RouteInfo) AfterDelete(tx *gorm.DB) (err error) {
	err = tx.Where("route_id = ?", classes.ID).Delete(&ClassesInfo{}).Error
	return err
}
