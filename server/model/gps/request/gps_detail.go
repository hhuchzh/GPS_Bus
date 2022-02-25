package request

import (
	//"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type GpsDetailSearch struct {
	GpsSn     string `json:"gpsSn" form:"gpsSn"`
	BeginTime string `json:"beginTime" form:"beginTime"`
	EndTime   string `json:"endTime" form:"endTime"`
	request.PageInfo
}

type BusDetailSearch struct {
	BusPlate  string `json:"busplate" form:"busplate"`
	BeginTime string `json:"beginTime" form:"beginTime"`
	EndTime   string `json:"endTime" form:"endTime"`
}
