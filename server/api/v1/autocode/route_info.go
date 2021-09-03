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

type RouteInfoApi struct {
}

var routeInfoService = service.ServiceGroupApp.AutoCodeServiceGroup.RouteInfoService


// CreateRouteInfo 创建RouteInfo
// @Tags RouteInfo
// @Summary 创建RouteInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.RouteInfo true "创建RouteInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /routeInfo/createRouteInfo [post]
func (routeInfoApi *RouteInfoApi) CreateRouteInfo(c *gin.Context) {
	var routeInfo autocode.RouteInfo
	_ = c.ShouldBindJSON(&routeInfo)
	if err := routeInfoService.CreateRouteInfo(routeInfo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRouteInfo 删除RouteInfo
// @Tags RouteInfo
// @Summary 删除RouteInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.RouteInfo true "删除RouteInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /routeInfo/deleteRouteInfo [delete]
func (routeInfoApi *RouteInfoApi) DeleteRouteInfo(c *gin.Context) {
	var routeInfo autocode.RouteInfo
	_ = c.ShouldBindJSON(&routeInfo)
	if err := routeInfoService.DeleteRouteInfo(routeInfo); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRouteInfoByIds 批量删除RouteInfo
// @Tags RouteInfo
// @Summary 批量删除RouteInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RouteInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /routeInfo/deleteRouteInfoByIds [delete]
func (routeInfoApi *RouteInfoApi) DeleteRouteInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := routeInfoService.DeleteRouteInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRouteInfo 更新RouteInfo
// @Tags RouteInfo
// @Summary 更新RouteInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.RouteInfo true "更新RouteInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /routeInfo/updateRouteInfo [put]
func (routeInfoApi *RouteInfoApi) UpdateRouteInfo(c *gin.Context) {
	var routeInfo autocode.RouteInfo
	_ = c.ShouldBindJSON(&routeInfo)
	if err := routeInfoService.UpdateRouteInfo(routeInfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRouteInfo 用id查询RouteInfo
// @Tags RouteInfo
// @Summary 用id查询RouteInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.RouteInfo true "用id查询RouteInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /routeInfo/findRouteInfo [get]
func (routeInfoApi *RouteInfoApi) FindRouteInfo(c *gin.Context) {
	var routeInfo autocode.RouteInfo
	_ = c.ShouldBindQuery(&routeInfo)
	if err, rerouteInfo := routeInfoService.GetRouteInfo(routeInfo.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerouteInfo": rerouteInfo}, c)
	}
}

// GetRouteInfoList 分页获取RouteInfo列表
// @Tags RouteInfo
// @Summary 分页获取RouteInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.RouteInfoSearch true "分页获取RouteInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /routeInfo/getRouteInfoList [get]
func (routeInfoApi *RouteInfoApi) GetRouteInfoList(c *gin.Context) {
	var pageInfo autocodeReq.RouteInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := routeInfoService.GetRouteInfoInfoList(pageInfo); err != nil {
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
