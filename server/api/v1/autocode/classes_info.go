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

type ClassesInfoApi struct {
}

var classesInfoService = service.ServiceGroupApp.AutoCodeServiceGroup.ClassesInfoService


// CreateClassesInfo 创建ClassesInfo
// @Tags ClassesInfo
// @Summary 创建ClassesInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ClassesInfo true "创建ClassesInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /classesInfo/createClassesInfo [post]
func (classesInfoApi *ClassesInfoApi) CreateClassesInfo(c *gin.Context) {
	var classesInfo autocode.ClassesInfo
	_ = c.ShouldBindJSON(&classesInfo)
	if err := classesInfoService.CreateClassesInfo(classesInfo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteClassesInfo 删除ClassesInfo
// @Tags ClassesInfo
// @Summary 删除ClassesInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ClassesInfo true "删除ClassesInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /classesInfo/deleteClassesInfo [delete]
func (classesInfoApi *ClassesInfoApi) DeleteClassesInfo(c *gin.Context) {
	var classesInfo autocode.ClassesInfo
	_ = c.ShouldBindJSON(&classesInfo)
	if err := classesInfoService.DeleteClassesInfo(classesInfo); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteClassesInfoByIds 批量删除ClassesInfo
// @Tags ClassesInfo
// @Summary 批量删除ClassesInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ClassesInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /classesInfo/deleteClassesInfoByIds [delete]
func (classesInfoApi *ClassesInfoApi) DeleteClassesInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := classesInfoService.DeleteClassesInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateClassesInfo 更新ClassesInfo
// @Tags ClassesInfo
// @Summary 更新ClassesInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ClassesInfo true "更新ClassesInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /classesInfo/updateClassesInfo [put]
func (classesInfoApi *ClassesInfoApi) UpdateClassesInfo(c *gin.Context) {
	var classesInfo autocode.ClassesInfo
	_ = c.ShouldBindJSON(&classesInfo)
	if err := classesInfoService.UpdateClassesInfo(classesInfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindClassesInfo 用id查询ClassesInfo
// @Tags ClassesInfo
// @Summary 用id查询ClassesInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.ClassesInfo true "用id查询ClassesInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /classesInfo/findClassesInfo [get]
func (classesInfoApi *ClassesInfoApi) FindClassesInfo(c *gin.Context) {
	var classesInfo autocode.ClassesInfo
	_ = c.ShouldBindQuery(&classesInfo)
	if err, reclassesInfo := classesInfoService.GetClassesInfo(classesInfo.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reclassesInfo": reclassesInfo}, c)
	}
}

// GetClassesInfoList 分页获取ClassesInfo列表
// @Tags ClassesInfo
// @Summary 分页获取ClassesInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.ClassesInfoSearch true "分页获取ClassesInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /classesInfo/getClassesInfoList [get]
func (classesInfoApi *ClassesInfoApi) GetClassesInfoList(c *gin.Context) {
	var pageInfo autocodeReq.ClassesInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := classesInfoService.GetClassesInfoInfoList(pageInfo); err != nil {
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
