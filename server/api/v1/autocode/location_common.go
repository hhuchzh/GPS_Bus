package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type LocationCommonApi struct {
}

var locationCommonService = service.ServiceGroupApp.AutoCodeServiceGroup.LocationCommonService


// CreateLocationCommon 创建LocationCommon
// @Tags LocationCommon
// @Summary 创建LocationCommon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.LocationCommon true "创建LocationCommon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /locationCommon/createLocationCommon [post]
func (locationCommonApi *LocationCommonApi) CreateLocationCommon(c *gin.Context) {
	var locationCommon autocode.LocationCommon
	_ = c.ShouldBindJSON(&locationCommon)
	if err := locationCommonService.CreateLocationCommon(locationCommon); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLocationCommon 删除LocationCommon
// @Tags LocationCommon
// @Summary 删除LocationCommon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.LocationCommon true "删除LocationCommon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /locationCommon/deleteLocationCommon [delete]
func (locationCommonApi *LocationCommonApi) DeleteLocationCommon(c *gin.Context) {
	var locationCommon autocode.LocationCommon
	_ = c.ShouldBindJSON(&locationCommon)
	if err := locationCommonService.DeleteLocationCommon(locationCommon); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteLocationCommonByIds 批量删除LocationCommon
// @Tags LocationCommon
// @Summary 批量删除LocationCommon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LocationCommon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /locationCommon/deleteLocationCommonByIds [delete]
func (locationCommonApi *LocationCommonApi) DeleteLocationCommonByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := locationCommonService.DeleteLocationCommonByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateLocationCommon 更新LocationCommon
// @Tags LocationCommon
// @Summary 更新LocationCommon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.LocationCommon true "更新LocationCommon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /locationCommon/updateLocationCommon [put]
func (locationCommonApi *LocationCommonApi) UpdateLocationCommon(c *gin.Context) {
	var locationCommon autocode.LocationCommon
	_ = c.ShouldBindJSON(&locationCommon)
	if err := locationCommonService.UpdateLocationCommon(locationCommon); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLocationCommon 用id查询LocationCommon
// @Tags LocationCommon
// @Summary 用id查询LocationCommon
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.LocationCommon true "用id查询LocationCommon"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /locationCommon/findLocationCommon [get]
func (locationCommonApi *LocationCommonApi) FindLocationCommon(c *gin.Context) {
	var locationCommon autocode.LocationCommon
	_ = c.ShouldBindQuery(&locationCommon)
	if err, relocationCommon := locationCommonService.GetLocationCommon(locationCommon.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relocationCommon": relocationCommon}, c)
	}
}

// GetLocationCommonList 分页获取LocationCommon列表
// @Tags LocationCommon
// @Summary 分页获取LocationCommon列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.LocationCommonSearch true "分页获取LocationCommon列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /locationCommon/getLocationCommonList [get]
func (locationCommonApi *LocationCommonApi) GetLocationCommonList(c *gin.Context) {
	var pageInfo autocodeReq.LocationCommonSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := locationCommonService.GetLocationCommonInfoList(pageInfo); err != nil {
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
