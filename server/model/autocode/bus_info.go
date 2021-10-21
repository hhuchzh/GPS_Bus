// 自动生成模板BusInfo
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

// BusInfo 结构体
// 如果含有time.Time 请自行import time包
type BusInfo struct {
	global.GVA_MODEL
	BusUserPhone string    `json:"busUserPhone" form:"busUserPhone" gorm:"column:bus_user_phone;comment:车主电话;type:varchar(100);"`
	BusUserName  string    `json:"busUserName" form:"busUserName" gorm:"column:bus_user_name;comment:车主姓名;type:varchar(100);"`
	BusPlate     string    `json:"busPlate" form:"busPlate" gorm:"column:bus_plate;comment:车牌;type:varchar(50);"`
	SeatNumber   *int      `json:"seatNumber" form:"seatNumber" gorm:"column:seat_number;comment:座位数;type:int"`
	CreateUserId *int      `json:"createUserId" form:"createUserId" gorm:"column:create_user_id;comment:创建用户id;type:bigint"`
	UpdateUserId *int      `json:"updateUserId" form:"updateUserId" gorm:"column:update_user_id;comment:更新用户id;type:bigint"`
	CompanyId    *int      `json:"companyId" form:"companyId" gorm:"column:company_id;comment:公司id;type:bigint"`
	GpsInfos     []GpsInfo `json:"gpsInfos" gorm:"foreignKey:BusId"`
}

// TableName BusInfo 表名
func (BusInfo) TableName() string {
	return "bus_info"
}

func (bus *BusInfo) AfterDelete(tx *gorm.DB) (err error) {
	return tx.Model(&GpsInfo{}).Where("bus_id = ?", bus.ID).Update("BusId", nil).Error
	return tx.Model(&ClassesInfo{}).Where("bus_id = ?", bus.ID).Update("BusId", 0).Error
}
