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

type ShiftInfoApi struct {
}

var shiftInfoService = service.ServiceGroupApp.AutoCodeServiceGroup.ShiftInfoService


// CreateShiftInfo 创建ShiftInfo
// @Tags ShiftInfo
// @Summary 创建ShiftInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ShiftInfo true "创建ShiftInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shiftInfo/createShiftInfo [post]
func (shiftInfoApi *ShiftInfoApi) CreateShiftInfo(c *gin.Context) {
	var shiftInfo autocode.ShiftInfo
	_ = c.ShouldBindJSON(&shiftInfo)
	if err := shiftInfoService.CreateShiftInfo(shiftInfo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteShiftInfo 删除ShiftInfo
// @Tags ShiftInfo
// @Summary 删除ShiftInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ShiftInfo true "删除ShiftInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shiftInfo/deleteShiftInfo [delete]
func (shiftInfoApi *ShiftInfoApi) DeleteShiftInfo(c *gin.Context) {
	var shiftInfo autocode.ShiftInfo
	_ = c.ShouldBindJSON(&shiftInfo)
	if err := shiftInfoService.DeleteShiftInfo(shiftInfo); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteShiftInfoByIds 批量删除ShiftInfo
// @Tags ShiftInfo
// @Summary 批量删除ShiftInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ShiftInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /shiftInfo/deleteShiftInfoByIds [delete]
func (shiftInfoApi *ShiftInfoApi) DeleteShiftInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := shiftInfoService.DeleteShiftInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateShiftInfo 更新ShiftInfo
// @Tags ShiftInfo
// @Summary 更新ShiftInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ShiftInfo true "更新ShiftInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shiftInfo/updateShiftInfo [put]
func (shiftInfoApi *ShiftInfoApi) UpdateShiftInfo(c *gin.Context) {
	var shiftInfo autocode.ShiftInfo
	_ = c.ShouldBindJSON(&shiftInfo)
	if err := shiftInfoService.UpdateShiftInfo(shiftInfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindShiftInfo 用id查询ShiftInfo
// @Tags ShiftInfo
// @Summary 用id查询ShiftInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.ShiftInfo true "用id查询ShiftInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shiftInfo/findShiftInfo [get]
func (shiftInfoApi *ShiftInfoApi) FindShiftInfo(c *gin.Context) {
	var shiftInfo autocode.ShiftInfo
	_ = c.ShouldBindQuery(&shiftInfo)
	if err, reshiftInfo := shiftInfoService.GetShiftInfo(shiftInfo.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reshiftInfo": reshiftInfo}, c)
	}
}

// GetShiftInfoList 分页获取ShiftInfo列表
// @Tags ShiftInfo
// @Summary 分页获取ShiftInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.ShiftInfoSearch true "分页获取ShiftInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shiftInfo/getShiftInfoList [get]
func (shiftInfoApi *ShiftInfoApi) GetShiftInfoList(c *gin.Context) {
	var pageInfo autocodeReq.ShiftInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := shiftInfoService.GetShiftInfoInfoList(pageInfo); err != nil {
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
