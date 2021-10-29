package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type MilesInfoService struct {
}

func (s *MilesInfoService) GetMilesInfoList(info autoCodeReq.MilesInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.MilesInfo{})
	var milesInfos []autocode.MilesInfo
	if info.ClassesId != nil {
		db = db.Where("`classes_id` = ?", info.ClassesId)
	}

	if info.CalcDate != "" {
		db = db.Where("`cal_date` = ?", info.CalcDate)
	}
	err = db.Count(&total).Error
	//err = db.Limit(limit).Offset(offset).Preload("Arrival").Find(&checkinInfos).Error
	err = db.Limit(limit).Offset(offset).Find(&milesInfos).Error
	return err, milesInfos, total
}
