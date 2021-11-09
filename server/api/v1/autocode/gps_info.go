package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GpsInfoApi struct {
}

var gpsInfoService = service.ServiceGroupApp.AutoCodeServiceGroup.GpsInfoService

// CreateGpsInfo 创建GpsInfo
// @Tags GpsInfo
// @Summary 创建GpsInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.GpsInfo true "创建GpsInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gpsInfo/createGpsInfo [post]
func (gpsInfoApi *GpsInfoApi) CreateGpsInfo(c *gin.Context) {
	var gpsInfo autocode.GpsInfo
	_ = c.ShouldBindJSON(&gpsInfo)
	if err := gpsInfoService.CreateGpsInfo(gpsInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteGpsInfo 删除GpsInfo
// @Tags GpsInfo
// @Summary 删除GpsInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.GpsInfo true "删除GpsInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gpsInfo/deleteGpsInfo [delete]
func (gpsInfoApi *GpsInfoApi) DeleteGpsInfo(c *gin.Context) {
	var gpsInfo autocode.GpsInfo
	_ = c.ShouldBindJSON(&gpsInfo)
	if err := gpsInfoService.DeleteGpsInfo(gpsInfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteGpsInfoByIds 批量删除GpsInfo
// @Tags GpsInfo
// @Summary 批量删除GpsInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除GpsInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /gpsInfo/deleteGpsInfoByIds [delete]
func (gpsInfoApi *GpsInfoApi) DeleteGpsInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := gpsInfoService.DeleteGpsInfoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateGpsInfo 更新GpsInfo
// @Tags GpsInfo
// @Summary 更新GpsInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.GpsInfo true "更新GpsInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /gpsInfo/updateGpsInfo [put]
func (gpsInfoApi *GpsInfoApi) UpdateGpsInfo(c *gin.Context) {
	var gpsInfo autocode.GpsInfo
	_ = c.ShouldBindJSON(&gpsInfo)
	if err := gpsInfoService.UpdateGpsInfo(gpsInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindGpsInfo 用id查询GpsInfo
// @Tags GpsInfo
// @Summary 用id查询GpsInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.GpsInfo true "用id查询GpsInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /gpsInfo/findGpsInfo [get]
func (gpsInfoApi *GpsInfoApi) FindGpsInfo(c *gin.Context) {
	var gpsInfo autocode.GpsInfo
	_ = c.ShouldBindQuery(&gpsInfo)
	if err, regpsInfo := gpsInfoService.GetGpsInfo(gpsInfo.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"regpsInfo": regpsInfo}, c)
	}
}

// GetGpsInfoList 分页获取GpsInfo列表
// @Tags GpsInfo
// @Summary 分页获取GpsInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.GpsInfoSearch true "分页获取GpsInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gpsInfo/getGpsInfoList [get]
func (gpsInfoApi *GpsInfoApi) GetGpsInfoList(c *gin.Context) {
	var pageInfo autocodeReq.GpsInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := gpsInfoService.GetGpsInfoInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetGpsInfoList 分页获取可用GpsInfo列表
// @Tags GpsInfo
// @Summary 分页获取可用GpsInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.GpsInfoSearch true "分页获取可用GpsInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gpsInfo/getAvailableGpsInfoList [get]
func (gpsInfoApi *GpsInfoApi) GetAvailableGpsInfoList(c *gin.Context) {
	var pageInfo autocodeReq.GpsInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := gpsInfoService.GetAvailableGpsInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (gpsInfoApi *GpsInfoApi) GetNotAvailableGpsInfoList(c *gin.Context) {
	var pageInfo autocodeReq.GpsInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := gpsInfoService.GetNotAvailableGpsInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
