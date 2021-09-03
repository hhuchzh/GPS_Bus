package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CheckinInfoRouter struct {
}

// InitCheckinInfoRouter 初始化 CheckinInfo 路由信息
func (s *CheckinInfoRouter) InitCheckinInfoRouter(Router *gin.RouterGroup) {
	checkinInfoRouter := Router.Group("checkinInfo").Use(middleware.OperationRecord())
	var checkinInfoApi = v1.ApiGroupApp.AutoCodeApiGroup.CheckinInfoApi
	{
		checkinInfoRouter.POST("createCheckinInfo", checkinInfoApi.CreateCheckinInfo)   // 新建CheckinInfo
		checkinInfoRouter.DELETE("deleteCheckinInfo", checkinInfoApi.DeleteCheckinInfo) // 删除CheckinInfo
		checkinInfoRouter.DELETE("deleteCheckinInfoByIds", checkinInfoApi.DeleteCheckinInfoByIds) // 批量删除CheckinInfo
		checkinInfoRouter.PUT("updateCheckinInfo", checkinInfoApi.UpdateCheckinInfo)    // 更新CheckinInfo
		checkinInfoRouter.GET("findCheckinInfo", checkinInfoApi.FindCheckinInfo)        // 根据ID获取CheckinInfo
		checkinInfoRouter.GET("getCheckinInfoList", checkinInfoApi.GetCheckinInfoList)  // 获取CheckinInfo列表
	}
}
