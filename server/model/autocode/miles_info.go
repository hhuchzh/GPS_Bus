package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type MilesInfo struct {
	global.GVA_MODEL
	CreateUserId *int         `json:"createUserId" form:"createUserId" gorm:"column:create_user_id;comment:创建人ID;type:bigint"`
	UpdateUserId *int         `json:"updateUserId" form:"updateUserId" gorm:"column:update_user_id;comment:修改人ID;type:bigint"`
	ClassesId    *uint        `json:"classesId" form:"classesId" gorm:"column:classes_id;comment:;type:bigint"`
	CalcDate string       `json:"calcDate" form:"calcDate" gorm:"column:cal_date;comment:;type:date"`
	Distance float64 `json:"distance" form:"distance" gorm:"column:distance;comment:;type:float"`
}

// TableName CheckinInfo 表名
func (MilesInfo) TableName() string {
	return "miles_info"
}
