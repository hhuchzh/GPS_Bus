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

type CheckinInfoApi struct {
}

var checkinInfoService = service.ServiceGroupApp.AutoCodeServiceGroup.CheckinInfoService


// CreateCheckinInfo 创建CheckinInfo
// @Tags CheckinInfo
// @Summary 创建CheckinInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CheckinInfo true "创建CheckinInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /checkinInfo/createCheckinInfo [post]
func (checkinInfoApi *CheckinInfoApi) CreateCheckinInfo(c *gin.Context) {
	var checkinInfo autocode.CheckinInfo
	_ = c.ShouldBindJSON(&checkinInfo)
	if err := checkinInfoService.CreateCheckinInfo(checkinInfo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCheckinInfo 删除CheckinInfo
// @Tags CheckinInfo
// @Summary 删除CheckinInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CheckinInfo true "删除CheckinInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /checkinInfo/deleteCheckinInfo [delete]
func (checkinInfoApi *CheckinInfoApi) DeleteCheckinInfo(c *gin.Context) {
	var checkinInfo autocode.CheckinInfo
	_ = c.ShouldBindJSON(&checkinInfo)
	if err := checkinInfoService.DeleteCheckinInfo(checkinInfo); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCheckinInfoByIds 批量删除CheckinInfo
// @Tags CheckinInfo
// @Summary 批量删除CheckinInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CheckinInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /checkinInfo/deleteCheckinInfoByIds [delete]
func (checkinInfoApi *CheckinInfoApi) DeleteCheckinInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := checkinInfoService.DeleteCheckinInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCheckinInfo 更新CheckinInfo
// @Tags CheckinInfo
// @Summary 更新CheckinInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CheckinInfo true "更新CheckinInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /checkinInfo/updateCheckinInfo [put]
func (checkinInfoApi *CheckinInfoApi) UpdateCheckinInfo(c *gin.Context) {
	var checkinInfo autocode.CheckinInfo
	_ = c.ShouldBindJSON(&checkinInfo)
	if err := checkinInfoService.UpdateCheckinInfo(checkinInfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCheckinInfo 用id查询CheckinInfo
// @Tags CheckinInfo
// @Summary 用id查询CheckinInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.CheckinInfo true "用id查询CheckinInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /checkinInfo/findCheckinInfo [get]
func (checkinInfoApi *CheckinInfoApi) FindCheckinInfo(c *gin.Context) {
	var checkinInfo autocode.CheckinInfo
	_ = c.ShouldBindQuery(&checkinInfo)
	if err, recheckinInfo := checkinInfoService.GetCheckinInfo(checkinInfo.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recheckinInfo": recheckinInfo}, c)
	}
}

// GetCheckinInfoList 分页获取CheckinInfo列表
// @Tags CheckinInfo
// @Summary 分页获取CheckinInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.CheckinInfoSearch true "分页获取CheckinInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /checkinInfo/getCheckinInfoList [get]
func (checkinInfoApi *CheckinInfoApi) GetCheckinInfoList(c *gin.Context) {
	var pageInfo autocodeReq.CheckinInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := checkinInfoService.GetCheckinInfoInfoList(pageInfo); err != nil {
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
