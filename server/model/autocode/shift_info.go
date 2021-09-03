// 自动生成模板ShiftInfo
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// ShiftInfo 结构体
// 如果含有time.Time 请自行import time包
type ShiftInfo struct {
	global.GVA_MODEL
	CreateUserId *int       `json:"createUserId" form:"createUserId" gorm:"column:create_user_id;comment:创建人ID;type:bigint"`
	UpdateUserId *int       `json:"updateUserId" form:"updateUserId" gorm:"column:update_user_id;comment:修改人ID;type:bigint"`
	BusId        *int       `json:"busId" form:"busId" gorm:"column:bus_id;comment:;type:bigint"`
	SendTime     *time.Time `json:"sendTime" form:"sendTime" gorm:"column:send_time;comment:;type:date"`
	ClassesId    *int       `json:"classesId" form:"classesId" gorm:"column:classes_id;comment:;type:bigint"`
	ShiftMiles   *int       `json:"shiftMiles" form:"shiftMiles" gorm:"column:shift_miles;comment:;type:int"`
}

// TableName ShiftInfo 表名
func (ShiftInfo) TableName() string {
	return "shift_info"
}
