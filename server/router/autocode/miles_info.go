package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MilesInfoRouter struct {
}

// InitCheckinInfoRouter 初始化 CheckinInfo 路由信息
func (s *MilesInfoRouter) InitMilesInfoRouter(Router *gin.RouterGroup) {
	milesInfoRouter := Router.Group("milesInfo").Use(middleware.OperationRecord())
	var milesInfoApi = v1.ApiGroupApp.AutoCodeApiGroup.MilesInfoApi
	{
		milesInfoRouter.GET("getMilesInfoList", milesInfoApi.GetMilesInfoList)
	}
}
