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

func (ga *GpsApi) ListLocation(c *gin.Context) {
	if location, err := gpsService.ListLocation(); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"location": location}, c)
	}
}

func (ga *GpsApi) GetLocation(c *gin.Context) {
    imei := c.Param("imei")
	if location, err := gpsService.GetLocation(imei); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"location": location}, c)
	}
}

func (ga *GpsApi) ListDevice(c *gin.Context) {
	if device, err := gpsService.ListDevice(); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"device": device}, c)
	}
}

func (ga *GpsApi) ListTrack(c *gin.Context) {
    imei := c.Param("imei")
    beginTime := c.Query("begin_time")
    endTime := c.Query("end_time")
    if beginTime == "" || endTime == "" {
		response.FailWithMessage("必须指定开始时间和结束时间", c)
        return
    }
	if track, err := gpsService.ListTrack(imei, beginTime, endTime); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"track": track}, c)
	}
}
