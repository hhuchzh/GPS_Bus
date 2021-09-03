// 自动生成模板GpsInfo
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// GpsInfo 结构体
// 如果含有time.Time 请自行import time包
type GpsInfo struct {
      global.GVA_MODEL
      CreateUserId  *int `json:"createUserId" form:"createUserId" gorm:"column:create_user_id;comment:创建人ID;type:bigint"`
      UpdateUserId  *int `json:"updateUserId" form:"updateUserId" gorm:"column:update_user_id;comment:修改人ID;type:bigint"`
      GpsSn  string `json:"gpsSn" form:"gpsSn" gorm:"column:gps_sn;comment:GPS序列号;type:varchar(100);"`
      GpsName  string `json:"gpsName" form:"gpsName" gorm:"column:gps_name;comment:GPS名字;type:varchar(100);"`
      BusId  *int `json:"busId" form:"busId" gorm:"column:bus_id;comment:;type:bigint"`
}


// TableName GpsInfo 表名
func (GpsInfo) TableName() string {
  return "gps_info"
}

