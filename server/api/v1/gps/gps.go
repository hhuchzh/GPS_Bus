package gps

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type GpsApi struct {
}

var gpsService = service.ServiceGroupApp.GpsServiceGroup.GpsService

// FindBusInfo 用id查询BusInfo
// @Tags BusInfo
// @Summary 用id查询BusInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.BusInfo true "用id查询BusInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /busInfo/findBusInfo [get]
func (ga *GpsApi) ListLocation(c *gin.Context) {
	if location, err := gpsService.ListLocation(); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"location": location}, c)
	}
}

