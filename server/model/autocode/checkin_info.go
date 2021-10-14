// 自动生成模板CheckinInfo
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CheckinInfo 结构体
// 如果含有time.Time 请自行import time包
type CheckinInfo struct {
	global.GVA_MODEL
	CreateUserId *int         `json:"createUserId" form:"createUserId" gorm:"column:create_user_id;comment:创建人ID;type:bigint"`
	UpdateUserId *int         `json:"updateUserId" form:"updateUserId" gorm:"column:update_user_id;comment:修改人ID;type:bigint"`
	ArrivalId    *int         `json:"locationId" form:"locationId" gorm:"column:location_id;comment:;type:bigint"`
	ClassesId    *uint64      `json:"classesId" form:"classesId" gorm:"column:classes_id;comment:;type:bigint"`
	CheckinTime  string       `json:"checkinTime" form:"checkinTime" gorm:"column:checkin_time;comment:;type:time"`
	CheckinDate  string       `json:"checkinDate" form:"checkinDate" gorm:"column:checkin_date;comment:;type:date"`
	Arrival      *ArrivalInfo `gorm:"foreignKey:ArrivalId"`
}

// TableName CheckinInfo 表名
func (CheckinInfo) TableName() string {
	return "checkin_info"
}
