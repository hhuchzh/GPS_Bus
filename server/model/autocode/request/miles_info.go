package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type MilesInfoSearch struct{
    autocode.MilesInfo
    request.PageInfo
}

type MilesInfoExport struct {
    Plate string `json:"plate" form:"plate"`
    BeginTime string `json:"beginTime" form:"beginTime"`
    EndTime string `json:"endTime" form:"endTime"`
}
