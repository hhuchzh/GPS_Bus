package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BusInfoRouter struct {
}

// InitBusInfoRouter 初始化 BusInfo 路由信息
func (s *BusInfoRouter) InitBusInfoRouter(Router *gin.RouterGroup) {
	busInfoRouter := Router.Group("busInfo").Use(middleware.OperationRecord())
	var busInfoApi = v1.ApiGroupApp.AutoCodeApiGroup.BusInfoApi
	{
		busInfoRouter.POST("createBusInfo", busInfoApi.CreateBusInfo)   // 新建BusInfo
		busInfoRouter.DELETE("deleteBusInfo", busInfoApi.DeleteBusInfo) // 删除BusInfo
		busInfoRouter.DELETE("deleteBusInfoByIds", busInfoApi.DeleteBusInfoByIds) // 批量删除BusInfo
		busInfoRouter.PUT("updateBusInfo", busInfoApi.UpdateBusInfo)    // 更新BusInfo
		busInfoRouter.GET("findBusInfo", busInfoApi.FindBusInfo)        // 根据ID获取BusInfo
		busInfoRouter.GET("getBusInfoList", busInfoApi.GetBusInfoList)  // 获取BusInfo列表
	}
}
