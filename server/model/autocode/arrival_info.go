// 自动生成模板ArrivalInfo
package autocode

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ArrivalInfo 结构体
// 如果含有time.Time 请自行import time包
type ArrivalInfo struct {
	global.GVA_MODEL
	CreateUserId *int       `json:"createUserId" form:"createUserId" gorm:"column:create_user_id;comment:创建人ID;type:bigint"`
	UpdateUserId *int       `json:"updateUserId" form:"updateUserId" gorm:"column:update_user_id;comment:修改人ID;type:bigint"`
	LocationId   *int       `json:"locationId" form:"locationId" gorm:"column:location_id;comment:;type:bigint"`
	ClassesId    *int       `json:"classesId" form:"classesId" gorm:"column:classes_id;comment:;type:bigint"`
	ArrivalTime  *time.Time `json:"arrivalTime" form:"arrivalTime" gorm:"column:arrival_time;comment:站点时间;type:time);"`
}

// TableName ArrivalInfo 表名
func (ArrivalInfo) TableName() string {
	return "arrival_info"
}
