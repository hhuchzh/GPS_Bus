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

type ArrivalInfoApi struct {
}

var arrivalInfoService = service.ServiceGroupApp.AutoCodeServiceGroup.ArrivalInfoService


// CreateArrivalInfo 创建ArrivalInfo
// @Tags ArrivalInfo
// @Summary 创建ArrivalInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ArrivalInfo true "创建ArrivalInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /arrivalInfo/createArrivalInfo [post]
func (arrivalInfoApi *ArrivalInfoApi) CreateArrivalInfo(c *gin.Context) {
	var arrivalInfo autocode.ArrivalInfo
	_ = c.ShouldBindJSON(&arrivalInfo)
	if err := arrivalInfoService.CreateArrivalInfo(arrivalInfo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteArrivalInfo 删除ArrivalInfo
// @Tags ArrivalInfo
// @Summary 删除ArrivalInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ArrivalInfo true "删除ArrivalInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /arrivalInfo/deleteArrivalInfo [delete]
func (arrivalInfoApi *ArrivalInfoApi) DeleteArrivalInfo(c *gin.Context) {
	var arrivalInfo autocode.ArrivalInfo
	_ = c.ShouldBindJSON(&arrivalInfo)
	if err := arrivalInfoService.DeleteArrivalInfo(arrivalInfo); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteArrivalInfoByIds 批量删除ArrivalInfo
// @Tags ArrivalInfo
// @Summary 批量删除ArrivalInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ArrivalInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /arrivalInfo/deleteArrivalInfoByIds [delete]
func (arrivalInfoApi *ArrivalInfoApi) DeleteArrivalInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := arrivalInfoService.DeleteArrivalInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateArrivalInfo 更新ArrivalInfo
// @Tags ArrivalInfo
// @Summary 更新ArrivalInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ArrivalInfo true "更新ArrivalInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /arrivalInfo/updateArrivalInfo [put]
func (arrivalInfoApi *ArrivalInfoApi) UpdateArrivalInfo(c *gin.Context) {
	var arrivalInfo autocode.ArrivalInfo
	_ = c.ShouldBindJSON(&arrivalInfo)
	if err := arrivalInfoService.UpdateArrivalInfo(arrivalInfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindArrivalInfo 用id查询ArrivalInfo
// @Tags ArrivalInfo
// @Summary 用id查询ArrivalInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.ArrivalInfo true "用id查询ArrivalInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /arrivalInfo/findArrivalInfo [get]
func (arrivalInfoApi *ArrivalInfoApi) FindArrivalInfo(c *gin.Context) {
	var arrivalInfo autocode.ArrivalInfo
	_ = c.ShouldBindQuery(&arrivalInfo)
	if err, rearrivalInfo := arrivalInfoService.GetArrivalInfo(arrivalInfo.ClassesId, arrivalInfo.LocationId); err != nil {
        //global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		//response.FailWithMessage("查询失败", c)
		response.FailWithData(gin.H{"rearrivalInfo": nil}, c)
	} else {
		response.OkWithData(gin.H{"rearrivalInfo": rearrivalInfo}, c)
	}
}

// GetArrivalInfoList 分页获取ArrivalInfo列表
// @Tags ArrivalInfo
// @Summary 分页获取ArrivalInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.ArrivalInfoSearch true "分页获取ArrivalInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /arrivalInfo/getArrivalInfoList [get]
func (arrivalInfoApi *ArrivalInfoApi) GetArrivalInfoList(c *gin.Context) {
	var pageInfo autocodeReq.ArrivalInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := arrivalInfoService.GetArrivalInfoInfoList(pageInfo); err != nil {
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
