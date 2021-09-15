package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type GpsDetailSearch struct{
    Imei string
    StartTime string
    EndTime string
    request.PageInfo
}
