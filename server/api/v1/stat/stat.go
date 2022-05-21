package stat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type StatApi struct {
}

var statService = service.ServiceGroupApp.StatServiceGroup.StatService

func (ga *StatApi) GetShuttleStatistics(c *gin.Context) {
	if stat, err := statService.GetShuttleStatistics(); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"stat": stat}, c)
	}
}
