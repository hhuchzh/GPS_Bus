// 自动生成模板LocationCommon
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// LocationCommon 结构体
// 如果含有time.Time 请自行import time包
type LocationCommon struct {
      global.GVA_MODEL
      CreateUserId  *int `json:"createUserId" form:"createUserId" gorm:"column:create_user_id;comment:创建人ID;type:bigint"`
      UpdateUserId  *int `json:"updateUserId" form:"updateUserId" gorm:"column:update_user_id;comment:修改人ID;type:bigint"`
      LocationName  string `json:"locationName" form:"locationName" gorm:"column:location_name;comment:站点名;type:varchar(100);"`
      Longtitude  string `json:"longtitude" form:"longtitude" gorm:"column:longtitude;comment:经度;type:varchar(50);"`
      Latitude  string `json:"latitude" form:"latitude" gorm:"column:latitude;comment:纬度;type:varchar(50);"`
}


// TableName LocationCommon 表名
func (LocationCommon) TableName() string {
  return "location_common"
}

