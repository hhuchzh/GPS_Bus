package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ShiftInfoRouter struct {
}

// InitShiftInfoRouter 初始化 ShiftInfo 路由信息
func (s *ShiftInfoRouter) InitShiftInfoRouter(Router *gin.RouterGroup) {
	shiftInfoRouter := Router.Group("shiftInfo").Use(middleware.OperationRecord())
	var shiftInfoApi = v1.ApiGroupApp.AutoCodeApiGroup.ShiftInfoApi
	{
		shiftInfoRouter.POST("createShiftInfo", shiftInfoApi.CreateShiftInfo)   // 新建ShiftInfo
		shiftInfoRouter.DELETE("deleteShiftInfo", shiftInfoApi.DeleteShiftInfo) // 删除ShiftInfo
		shiftInfoRouter.DELETE("deleteShiftInfoByIds", shiftInfoApi.DeleteShiftInfoByIds) // 批量删除ShiftInfo
		shiftInfoRouter.PUT("updateShiftInfo", shiftInfoApi.UpdateShiftInfo)    // 更新ShiftInfo
		shiftInfoRouter.GET("findShiftInfo", shiftInfoApi.FindShiftInfo)        // 根据ID获取ShiftInfo
		shiftInfoRouter.GET("getShiftInfoList", shiftInfoApi.GetShiftInfoList)  // 获取ShiftInfo列表
	}
}
