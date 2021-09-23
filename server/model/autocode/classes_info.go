// 自动生成模板ClassesInfo
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

// ClassesInfo 结构体
// 如果含有time.Time 请自行import time包
type ClassesInfo struct {
	global.GVA_MODEL
	CreateUserId *int     `json:"createUserId" form:"createUserId" gorm:"column:create_user_id;comment:创建人ID;type:bigint"`
	UpdateUserId *int     `json:"updateUserId" form:"updateUserId" gorm:"column:update_user_id;comment:修改人ID;type:bigint"`
	RouteId      *uint    `json:"routeId" form:"routeId" gorm:"column:route_id;comment:;type:bigint"`
	Remark       string   `json:"remark" form:"remark" gorm:"column:remark;comment:备注;type:varchar(100);"`
	ClassesTime  string   `json:"classesTime" form:"classesTime" gorm:"column:classes_time;comment:发车时间;type:time);"`
	BusId        *uint    `json:"busId" form:"busId" gorm:"column:bus_id;comment:;type:bigint"`
	Bus          *BusInfo `json:"busInfo" form:"busInfo" gorm:"foreignKey:BusId"`
}

// TableName ClassesInfo 表名
func (ClassesInfo) TableName() string {
	return "classes_info"
}

// func (classes *ClassesInfo) AfterCreate(tx *gorm.DB) (err error) {
// 	locations := make([]LocationInfo, 0)
// 	err = tx.Where("route_id = ?", *classes.RouteId).Find(locations).Error
// 	if err != nil {
// 		return err
// 	}

// 	for _, location := range locations {
// 		arrival := ArrivalInfo{LocationId: &location.ID, ClassesId: &classes.ID}
// 		err := tx.Create(&arrival).Error
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

func (classes *ClassesInfo) AfterDelete(tx *gorm.DB) (err error) {
	err = tx.Where("classes_id = ?", classes.ID).Delete(&ArrivalInfo{}).Error
	return err
}
