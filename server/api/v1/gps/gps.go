package gps

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	gpsreq "github.com/flipped-aurora/gin-vue-admin/server/model/gps/request"
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
	var gs gpsreq.GpsDetailSearch
	_ = c.ShouldBindQuery(&gs)
	if location, err := gpsService.GetLocation(gs.GpsSn); err != nil {
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
	var gs gpsreq.GpsDetailSearch
	_ = c.ShouldBindQuery(&gs)

	global.GVA_LOG.Info("gps ...", zap.String("GpsSn", gs.GpsSn))

	if gs.BeginTime == "" || gs.EndTime == "" {
		response.FailWithMessage("必须指定开始时间和结束时间", c)
		return
	}
	if track, err := gpsService.ListTrack(gs.GpsSn, gs.BeginTime, gs.EndTime); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"track": track}, c)
	}
}

func (ga *GpsApi) ListBusTrack(c *gin.Context) {
	var req gpsreq.BusDetailSearch
	if err := c.ShouldBindQuery(&req); err != nil {
		global.GVA_LOG.Error("转换失败!", zap.Any("err", err))
	}
	
	fmt.Println("%+v", *c)

	global.GVA_LOG.Info("busplate ...", zap.String("BusPlate", req.BusPlate))
	global.GVA_LOG.Info("begintime ...", zap.String("begin", req.BeginTime))

	if req.BusPlate == "" {
		response.FailWithMessage("必须指定车牌号", c)
		return
	}

	if req.BeginTime == "" || req.EndTime == "" {
		response.FailWithMessage("必须指定开始时间和结束时间", c)
		return
	}

	var businfo autocode.BusInfo
	res := global.GVA_DB.Where("bus_plate = ?", req.BusPlate).Find(&businfo)
	var cnt int64
	_ = res.Count(&cnt)
	if cnt == 0 {
		response.FailWithMessage("车牌号不正确", c)
		return
	}

	cnt = 0
	var gps autocode.GpsInfo
	result := global.GVA_DB.Where("bus_id = ?", businfo.ID).Find(&gps)
	_ = result.Count(&cnt)
	if cnt == 0 {
		response.FailWithMessage("找不到对应的GPS设备", c)
		return
	}

	if track, err := gpsService.ListTrack(gps.GpsSn, req.BeginTime, req.EndTime); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"track": track}, c)
	}

}
