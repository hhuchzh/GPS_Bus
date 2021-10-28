package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type MilesInfoApi struct {
}

var milesInfoService = service.ServiceGroupApp.AutoCodeServiceGroup.MilesInfoService


func (milesInfoApi *MilesInfoApi) GetMilesInfoList(c *gin.Context) {
	var pageInfo autocodeReq.MilesInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := milesInfoService.GetMilesInfoList(pageInfo); err != nil {
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
