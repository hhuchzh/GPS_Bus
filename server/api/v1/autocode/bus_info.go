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

type BusInfoApi struct {
}

var busInfoService = service.ServiceGroupApp.AutoCodeServiceGroup.BusInfoService


// CreateBusInfo 创建BusInfo
// @Tags BusInfo
// @Summary 创建BusInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.BusInfo true "创建BusInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busInfo/createBusInfo [post]
func (busInfoApi *BusInfoApi) CreateBusInfo(c *gin.Context) {
	var busInfo autocode.BusInfo
	_ = c.ShouldBindJSON(&busInfo)
	if err := busInfoService.CreateBusInfo(busInfo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBusInfo 删除BusInfo
// @Tags BusInfo
// @Summary 删除BusInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.BusInfo true "删除BusInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /busInfo/deleteBusInfo [delete]
func (busInfoApi *BusInfoApi) DeleteBusInfo(c *gin.Context) {
	var busInfo autocode.BusInfo
	_ = c.ShouldBindJSON(&busInfo)
	if err := busInfoService.DeleteBusInfo(busInfo); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBusInfoByIds 批量删除BusInfo
// @Tags BusInfo
// @Summary 批量删除BusInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BusInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /busInfo/deleteBusInfoByIds [delete]
func (busInfoApi *BusInfoApi) DeleteBusInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := busInfoService.DeleteBusInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBusInfo 更新BusInfo
// @Tags BusInfo
// @Summary 更新BusInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.BusInfo true "更新BusInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /busInfo/updateBusInfo [put]
func (busInfoApi *BusInfoApi) UpdateBusInfo(c *gin.Context) {
	var busInfo autocode.BusInfo
	_ = c.ShouldBindJSON(&busInfo)
	if err := busInfoService.UpdateBusInfo(busInfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBusInfo 用id查询BusInfo
// @Tags BusInfo
// @Summary 用id查询BusInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.BusInfo true "用id查询BusInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /busInfo/findBusInfo [get]
func (busInfoApi *BusInfoApi) FindBusInfo(c *gin.Context) {
	var busInfo autocode.BusInfo
	_ = c.ShouldBindQuery(&busInfo)
	if err, rebusInfo := busInfoService.GetBusInfo(busInfo.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebusInfo": rebusInfo}, c)
	}
}

// GetBusInfoList 分页获取BusInfo列表
// @Tags BusInfo
// @Summary 分页获取BusInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.BusInfoSearch true "分页获取BusInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busInfo/getBusInfoList [get]
func (busInfoApi *BusInfoApi) GetBusInfoList(c *gin.Context) {
	var pageInfo autocodeReq.BusInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := busInfoService.GetBusInfoInfoList(pageInfo); err != nil {
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
