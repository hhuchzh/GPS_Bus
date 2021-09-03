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

type LocationInfoApi struct {
}

var locationInfoService = service.ServiceGroupApp.AutoCodeServiceGroup.LocationInfoService


// CreateLocationInfo 创建LocationInfo
// @Tags LocationInfo
// @Summary 创建LocationInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.LocationInfo true "创建LocationInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /locationInfo/createLocationInfo [post]
func (locationInfoApi *LocationInfoApi) CreateLocationInfo(c *gin.Context) {
	var locationInfo autocode.LocationInfo
	_ = c.ShouldBindJSON(&locationInfo)
	if err := locationInfoService.CreateLocationInfo(locationInfo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLocationInfo 删除LocationInfo
// @Tags LocationInfo
// @Summary 删除LocationInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.LocationInfo true "删除LocationInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /locationInfo/deleteLocationInfo [delete]
func (locationInfoApi *LocationInfoApi) DeleteLocationInfo(c *gin.Context) {
	var locationInfo autocode.LocationInfo
	_ = c.ShouldBindJSON(&locationInfo)
	if err := locationInfoService.DeleteLocationInfo(locationInfo); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteLocationInfoByIds 批量删除LocationInfo
// @Tags LocationInfo
// @Summary 批量删除LocationInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LocationInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /locationInfo/deleteLocationInfoByIds [delete]
func (locationInfoApi *LocationInfoApi) DeleteLocationInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := locationInfoService.DeleteLocationInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateLocationInfo 更新LocationInfo
// @Tags LocationInfo
// @Summary 更新LocationInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.LocationInfo true "更新LocationInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /locationInfo/updateLocationInfo [put]
func (locationInfoApi *LocationInfoApi) UpdateLocationInfo(c *gin.Context) {
	var locationInfo autocode.LocationInfo
	_ = c.ShouldBindJSON(&locationInfo)
	if err := locationInfoService.UpdateLocationInfo(locationInfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLocationInfo 用id查询LocationInfo
// @Tags LocationInfo
// @Summary 用id查询LocationInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.LocationInfo true "用id查询LocationInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /locationInfo/findLocationInfo [get]
func (locationInfoApi *LocationInfoApi) FindLocationInfo(c *gin.Context) {
	var locationInfo autocode.LocationInfo
	_ = c.ShouldBindQuery(&locationInfo)
	if err, relocationInfo := locationInfoService.GetLocationInfo(locationInfo.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relocationInfo": relocationInfo}, c)
	}
}

// GetLocationInfoList 分页获取LocationInfo列表
// @Tags LocationInfo
// @Summary 分页获取LocationInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.LocationInfoSearch true "分页获取LocationInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /locationInfo/getLocationInfoList [get]
func (locationInfoApi *LocationInfoApi) GetLocationInfoList(c *gin.Context) {
	var pageInfo autocodeReq.LocationInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := locationInfoService.GetLocationInfoInfoList(pageInfo); err != nil {
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
